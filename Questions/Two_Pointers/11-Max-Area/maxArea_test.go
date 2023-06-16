package arrays

import "testing"

func TestMaxArea(t *testing.T) {
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	height2 := []int{2, 2}

	t.Run("Test long buckets", func(t *testing.T) {
		want := 49

		got := maxArea(heights)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test short buckets", func(t *testing.T) {
		want := 2

		got := maxArea(height2)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
