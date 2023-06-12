package arrays

import "testing"

func TestCandy(t *testing.T) {
	ratings := []int{1, 2, 3, 3, 3, 2, 1}

	t.Run("Standard test case", func(t *testing.T) {
		want := 13

		got := candy(ratings)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
