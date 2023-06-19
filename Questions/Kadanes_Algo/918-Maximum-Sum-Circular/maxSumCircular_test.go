package kadanes

import "testing"

func TestMaxSubarraySumCircular(t *testing.T) {
	single := []int{1}
	real := []int{1, -2, 3, -2}

	t.Run("Test single array", func(t *testing.T) {
		want := 1

		got := maxSubarraySumCircular(single)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test real array", func(t *testing.T) {
		want := 3

		got := maxSubarraySumCircular(real)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
