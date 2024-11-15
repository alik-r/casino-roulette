package roulette

import (
	"testing"
)

func TestSpin(t *testing.T) {
	result := Spin()
	if result.Number < 0 || result.Number > 36 {
		t.Errorf("Invalid number: got %d, want between 0 and 36", result.Number)
	}
	if result.Color != numberColorMap[result.Number] {
		t.Errorf("Invalid color: got %s, want %s", result.Color, numberColorMap[result.Number])
	}
}

func TestPayout(t *testing.T) {
	tests := []struct {
		betType  BetType
		betValue interface{}
		result   RouletteResult
		expected int
	}{
		{ColorBet, "red", RouletteResult{Red, 1}, 1},
		{ColorBet, "black", RouletteResult{Black, 2}, 1},
		{ColorBet, "green", RouletteResult{Green, 0}, 35},
		{EvenOddBet, "even", RouletteResult{Red, 2}, 1},
		{EvenOddBet, "odd", RouletteResult{Black, 3}, 1},
		{HighLowBet, "high", RouletteResult{Red, 20}, 1},
		{HighLowBet, "low", RouletteResult{Black, 10}, 1},
		{NumberBet, 7.0, RouletteResult{Red, 7}, 35},
		{NumberBet, 8.0, RouletteResult{Black, 9}, 0},
	}

	for _, tt := range tests {
		t.Run(string(tt.betType), func(t *testing.T) {
			actual := Payout(tt.betType, tt.betValue, tt.result)
			if actual != tt.expected {
				t.Errorf("Payout(%v, %v, %v) = %d; want %d", tt.betType, tt.betValue, tt.result, actual, tt.expected)
			}
		})
	}
}

func TestIsValidBet(t *testing.T) {
	tests := []struct {
		betType  BetType
		betValue interface{}
		expected bool
	}{
		{ColorBet, "red", true},
		{ColorBet, "blue", false},
		{EvenOddBet, "even", true},
		{EvenOddBet, "odd", true},
		{EvenOddBet, "none", false},
		{HighLowBet, "high", true},
		{HighLowBet, "low", true},
		{HighLowBet, "medium", false},
		{NumberBet, 7.0, true},
		{NumberBet, 37.0, false},
		{NumberBet, -1.0, false},
	}

	for _, tt := range tests {
		t.Run(string(tt.betType), func(t *testing.T) {
			actual := IsValidBet(tt.betType, tt.betValue)
			if actual != tt.expected {
				t.Errorf("IsValidBet(%v, %v) = %v; want %v", tt.betType, tt.betValue, actual, tt.expected)
			}
		})
	}
}
