package arrays

import "testing"

func TesthIndex(t *testing.T) {
	nums1 := []int{3, 0, 6, 1, 5}

	t.Run("Standard test", func(t *testing.T) {
		want := 3

		got := hIndex(nums1)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
