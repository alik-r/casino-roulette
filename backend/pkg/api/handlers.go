package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/alik-r/casino-roulette/backend/pkg/auth"
	"github.com/alik-r/casino-roulette/backend/pkg/db"
	"github.com/alik-r/casino-roulette/backend/pkg/middleware"
	"github.com/alik-r/casino-roulette/backend/pkg/models"
	"github.com/alik-r/casino-roulette/backend/pkg/roulette"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   string `json:"avatar,omitempty"`
}

type DepositRequest struct {
	Amount int `json:"amount"`
}

type BetRequest struct {
	BetAmount int         `json:"bet_amount"`
	BetType   string      `json:"bet_type"`
	BetValue  interface{} `json:"bet_value"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	isRegister := false
	query := db.DB.Where("email = ?", loginRequest.Email)
	if loginRequest.Username != "" {
		query = query.Or("username = ?", loginRequest.Username)
	}

	err := query.First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if loginRequest.Username == "" || loginRequest.Email == "" {
				http.Error(w, "Both username and email are required", http.StatusBadRequest)
				return
			}

			if loginRequest.Avatar == "" {
				loginRequest.Avatar = "images/avatars/avatar1.png"
			} else {
				loginRequest.Avatar = "images/avatars/" + strings.Split(loginRequest.Avatar, "images/avatars/")[1]
			}

			user = models.User{
				Username: loginRequest.Username,
				Email:    loginRequest.Email,
				Avatar:   loginRequest.Avatar,
				Password: loginRequest.Password,
				Balance:  1000,
			}
			if err := db.DB.Create(&user).Error; err != nil {
				http.Error(w, "Failed to create user", http.StatusInternalServerError)
				return
			}
			isRegister = true
		} else {
			http.Error(w, "Error checking user", http.StatusInternalServerError)
			return
		}
	} else {
		if user.Password != loginRequest.Password {
			http.Error(w, "Invalid password", http.StatusBadRequest)
			return
		}
	}

	token, err := auth.GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token":       token,
		"is_register": strconv.FormatBool(isRegister),
	})
}

func Deposit(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value(middleware.UsernameKey).(string)
	if !ok || username == "" {
		http.Error(w, "Unauthenticated user", http.StatusUnauthorized)
		return
	}

	var deposit DepositRequest
	err := json.NewDecoder(r.Body).Decode(&deposit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if deposit.Amount <= 0 {
		http.Error(w, "Invalid deposit amount", http.StatusBadRequest)
		return
	}

	var user models.User
	err = db.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		user.Balance += deposit.Amount
		if err := db.DB.Save(&user).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func PlaceBet(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value(middleware.UsernameKey).(string)
	if !ok || username == "" {
		http.Error(w, "Unauthenticated user", http.StatusUnauthorized)
		return
	}

	var betRequest BetRequest
	err := json.NewDecoder(r.Body).Decode(&betRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if betRequest.BetAmount <= 0 {
		http.Error(w, "Bet amount must be greater than zero", http.StatusBadRequest)
		return
	}

	if !roulette.IsValidBet(roulette.BetType(betRequest.BetType), betRequest.BetValue) {
		log.Println("Invalid bet type or value:", betRequest.BetType, roulette.BetType(betRequest.BetType), betRequest.BetValue)
		http.Error(w, "Invalid bet type or value", http.StatusBadRequest)
		return
	}

	var user models.User
	err = db.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		log.Println("Database error in PlaceBet:", err)
		return
	}

	if user.Balance < betRequest.BetAmount {
		http.Error(w, "Insufficient balance", http.StatusBadRequest)
		return
	}

	result := roulette.Spin()

	payoutMultiplier := roulette.Payout(roulette.BetType(betRequest.BetType), betRequest.BetValue, result)
	var payout int
	var betResult string
	if payoutMultiplier > 0 {
		payout = betRequest.BetAmount * payoutMultiplier
		user.Balance += payout
		betResult = "win"
	} else {
		payout = 0
		user.Balance -= betRequest.BetAmount
		betResult = "lose"
	}

	if err := db.DB.Save(&user).Error; err != nil {
		http.Error(w, "Failed to update user balance", http.StatusInternalServerError)
		return
	}

	var betValueStr string
	switch betRequest.BetType {
	case "color", "evenodd", "highlow":
		value, ok := betRequest.BetValue.(string)
		if !ok {
			log.Printf("Invalid bet value, expected string, got %T", betRequest.BetValue)
			http.Error(w, "Invalid bet value", http.StatusBadRequest)
			return
		}
		betValueStr = value
	case "number":
		value, ok := betRequest.BetValue.(float64)
		if !ok {
			log.Printf("Invalid bet value, expected float64, got %T", betRequest.BetValue)
			http.Error(w, "Invalid bet value", http.StatusBadRequest)
			return
		}
		betValueStr = strconv.Itoa(int(value))
	default:
		log.Println("Invalid bet type:", betRequest.BetType)
		http.Error(w, "Invalid bet type", http.StatusBadRequest)
		return
	}

	bet := models.Bet{
		UserID:    user.ID,
		BetAmount: betRequest.BetAmount,
		BetType:   betRequest.BetType,
		BetValue:  betValueStr,
		Payout:    payout,
		Result:    betResult,
	}

	if err := db.DB.Create(&bet).Error; err != nil {
		http.Error(w, "Failed to record bet", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"balance":       user.Balance,
		"bet_amount":    bet.BetAmount,
		"bet_type":      bet.BetType,
		"bet_value":     bet.BetValue,
		"payout":        bet.Payout,
		"result":        bet.Result,
		"result_color":  result.Color,
		"result_number": result.Number,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value(middleware.UsernameKey).(string)
	if !ok || username == "" {
		http.Error(w, "Unauthenticated user", http.StatusUnauthorized)
		return
	}

	var users []models.User
	if err := db.DB.Order("balance DESC").Limit(10).Find(&users).Error; err != nil {
		http.Error(w, "Failed to retrieve leaderboard", http.StatusInternalServerError)
		return
	}

	type LeaderboardRow struct {
		Username string `json:"username"`
		Balance  int    `json:"balance"`
		Avatar   string `json:"avatar"`
	}

	var leaderboard []LeaderboardRow
	for _, user := range users {
		leaderboard = append(leaderboard, LeaderboardRow{
			Username: user.Username,
			Balance:  user.Balance,
			Avatar:   user.Avatar,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(leaderboard)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value(middleware.UsernameKey).(string)
	if !ok || username == "" {
		http.Error(w, "Unauthenticated user", http.StatusUnauthorized)
		return
	}

	var user models.User
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	response := map[string]interface{}{
		"username": user.Username,
		"email":    user.Email,
		"avatar":   user.Avatar,
		"balance":  user.Balance,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
