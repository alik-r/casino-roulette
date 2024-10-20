package api

import (
	"encoding/json"
	"net/http"

	"github.com/alik-r/casino-roulette/backend/pkg/db"
	"github.com/alik-r/casino-roulette/backend/pkg/models"
	"github.com/alik-r/casino-roulette/backend/pkg/roulette"
	"gorm.io/gorm"
)

type DepositRequest struct {
	Username string `json:"username"`
	Amount   int    `json:"amount"`
}

type BetRequest struct {
	Username  string `json:"username"`
	BetAmount int    `json:"bet_amount"`
	BetColor  string `json:"bet_color"`
}

func RegisterOrDeposit(w http.ResponseWriter, r *http.Request) {
	var deposit DepositRequest
	err := json.NewDecoder(r.Body).Decode(&deposit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if deposit.Username == "" || deposit.Amount <= 0 {
		http.Error(w, "invalid username or amount", http.StatusBadRequest)
		return
	}

	var user models.User
	err = db.DB.Where("username = ?", deposit.Username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			user = models.User{
				Username: deposit.Username,
				Balance:  deposit.Amount,
			}
			if err := db.DB.Create(&user).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
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
	var betRequest BetRequest
	err := json.NewDecoder(r.Body).Decode(&betRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if betRequest.Username == "" || betRequest.BetAmount <= 0 || !roulette.IsValidColor(betRequest.BetColor) {
		http.Error(w, "invalid username, bet amount or color", http.StatusBadRequest)
		return
	}

	var user models.User
	err = db.DB.Where("username = ?", betRequest.Username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.Balance < betRequest.BetAmount {
		http.Error(w, "insufficient balance", http.StatusBadRequest)
		return
	}

	result := roulette.Spin()
	payout := roulette.Payout(betRequest.BetColor, string(result.Color))
	var betResult string
	if payout > 0 {
		user.Balance += betRequest.BetAmount * (payout - 1)
		betResult = "win"
	} else {
		user.Balance -= betRequest.BetAmount
		betResult = "lose"
	}

	if err := db.DB.Save(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bet := models.Bet{
		UserID:    user.ID,
		BetAmount: betRequest.BetAmount,
		BetColor:  betRequest.BetColor,
		Result:    betResult,
	}

	if err := db.DB.Create(&bet).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"username":     user.Username,
		"balance":      user.Balance,
		"bet_amount":   bet.BetAmount,
		"bet_color":    bet.BetColor,
		"result":       bet.Result,
		"result_color": result.Color,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func GetLeaderBoard(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	err := db.DB.Order("balance desc").Find(&users).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
