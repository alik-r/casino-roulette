package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Avatar    string    `json:"avatar"`
	Password  string    `gorm:"not null" json:"-"`
	Balance   int       `gorm:"default:0" json:"balance"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	Bets      []Bet     `json:"bets,omitempty"`
}

type Bet struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	BetAmount int       `json:"bet_amount"`
	BetType   string    `json:"bet_type"`
	BetValue  string    `json:"bet_value"`
	Payout    int       `json:"payout"`
	Result    string    `json:"result"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
}
