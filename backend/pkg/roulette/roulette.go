package roulette

import (
	"time"

	"golang.org/x/exp/rand"
)

type Color string

const (
	Red   Color = "red"
	Black Color = "black"
	Green Color = "green"
)

type RouletteResult struct {
	Color Color `json:"color"`
}

func Spin() RouletteResult {
	rand.Seed(uint64(time.Now().UnixNano()))
	num := rand.Intn(100) + 1 // 1-100

	switch {
	case num <= 48:
		return RouletteResult{Color: Red}
	case num <= 96:
		return RouletteResult{Color: Black}
	default:
		return RouletteResult{Color: Green}
	}
}

func Payout(betColor, resultColor string) int {
	if betColor == resultColor {
		if betColor == string(Green) {
			return 14
		}
		return 2
	}
	return 0
}

func IsValidColor(color string) bool {
	switch color {
	case string(Red), string(Black), string(Green):
		return true
	default:
		return false
	}
}
