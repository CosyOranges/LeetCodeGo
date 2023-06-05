package arrays

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}

	t.Run("Testing normal case", func(t *testing.T) {
		want := 5

		got := maxProfit(prices)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func TestMaxProfitII(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}

	t.Run("Testing normal case", func(t *testing.T) {
		want := 7

		got := maxProfitII(prices)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
