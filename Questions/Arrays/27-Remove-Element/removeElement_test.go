package arrays

import (
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 3, 3, 3}
	nums2 := []int{4, 5}
	nums3 := []int{1}

	t.Run("Test regular array", func(t *testing.T) {
		want := 5

		got := removeElement(nums1, 2)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test with an array of length 2", func(t *testing.T) {
		want := 1

		got := removeElement(nums2, 5)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test with an array of length 1", func(t *testing.T) {
		want := 0

		got := removeElement(nums3, 1)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test with array that doesn't have the val", func(t *testing.T) {
		want := 6

		got := removeElement(nums1, -1)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
