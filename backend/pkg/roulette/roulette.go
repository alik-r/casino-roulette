package roulette

import (
	"log"
	"time"

	"golang.org/x/exp/rand"
)

type Color string
type BetType string

const (
	Red   Color = "red"
	Black Color = "black"
	Green Color = "green"
)

const (
	ColorBet   BetType = "color"
	EvenOddBet BetType = "evenodd"
	HighLowBet BetType = "highlow"
	NumberBet  BetType = "number"
)

type RouletteResult struct {
	Color  Color `json:"color"`
	Number int   `json:"number"`
}

var numberColorMap = map[int]Color{
	0: Green,
	1: Red, 2: Black, 3: Red, 4: Black, 5: Red, 6: Black, 7: Red, 8: Black, 9: Red, 10: Black,
	11: Black, 12: Red, 13: Black, 14: Red, 15: Black, 16: Red, 17: Black, 18: Red,
	19: Red, 20: Black, 21: Red, 22: Black, 23: Red, 24: Black, 25: Red, 26: Black, 27: Red, 28: Black,
	29: Black, 30: Red, 31: Black, 32: Red, 33: Black, 34: Red, 35: Black, 36: Red,
}

func Spin() RouletteResult {
	rand.Seed(uint64(time.Now().UnixNano()))
	num := rand.Intn(37) // 0-36
	color, exists := numberColorMap[num]
	if !exists {
		color = Green
	}

	return RouletteResult{
		Color:  color,
		Number: num,
	}
}

func Payout(betType BetType, betValue interface{}, result RouletteResult) int {
	switch betType {
	case ColorBet:
		betColor := betValue.(string)
		if betColor == string(result.Color) {
			if betColor == string(Green) {
				return 35
			}
			return 1
		}
	case EvenOddBet:
		betParity := betValue.(string)
		if betParity == "even" && result.Number%2 == 0 {
			return 1
		}
		if betParity == "odd" && result.Number%2 == 1 {
			return 1
		}
	case HighLowBet:
		betRange := betValue.(string)
		if betRange == "high" && result.Number >= 19 {
			return 1
		}
		if betRange == "low" && result.Number <= 18 {
			return 1
		}
	case NumberBet:
		number := betValue.(float64)
		betNumber := int(number)
		if betNumber == result.Number {
			return 35
		}
	}
	return 0
}

func IsValidBet(betType BetType, betValue interface{}) bool {
	switch betType {
	case ColorBet:
		color, ok := betValue.(string)
		if !ok {
			return false
		}
		return color == string(Red) || color == string(Black) || color == string(Green)
	case EvenOddBet:
		parity, ok := betValue.(string)
		if !ok {
			return false
		}
		return parity == "even" || parity == "odd"
	case HighLowBet:
		rangeValue, ok := betValue.(string)
		if !ok {
			return false
		}
		return rangeValue == "high" || rangeValue == "low"
	case NumberBet:
		numFloat, ok := betValue.(float64)
		if !ok {
			log.Println("number bet value is not int")
			return false
		}
		num := int(numFloat)
		return num >= 0 && num <= 36
	default:
		log.Println("invalid bet type", betType)
		return false
	}
}
