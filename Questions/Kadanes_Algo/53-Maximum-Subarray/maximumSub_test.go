package kadanes

import "testing"

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	empt := []int{}

	t.Run("Test standard array", func(t *testing.T) {
		want := 6

		got := maxSubArray(nums)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test empty array", func(t *testing.T) {
		want := 0

		got := maxSubArray(empt)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
