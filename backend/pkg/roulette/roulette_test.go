package roulette

import (
	"testing"
)

func TestSpin(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := Spin()
		if result.Color != Red && result.Color != Black && result.Color != Green {
			t.Errorf("Invalid result color: %v", result.Color)
		}
	}
}

func TestPayout(t *testing.T) {
	tests := []struct {
		betColor    string
		resultColor string
		expected    int
	}{
		{string(Red), string(Red), 2},
		{string(Black), string(Black), 2},
		{string(Green), string(Green), 14},
		{string(Red), string(Black), 0},
		{string(Black), string(Red), 0},
		{string(Green), string(Red), 0},
	}

	for _, tt := range tests {
		got := Payout(tt.betColor, tt.resultColor)
		if got != tt.expected {
			t.Errorf("Payout(%q, %q) = %d; want %d", tt.betColor, tt.resultColor, got, tt.expected)
		}
	}
}
