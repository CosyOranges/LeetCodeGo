package arrays

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	nums1 := []int{1, 2, 2, 3, 3, 4, 4, 5, 5}
	nums2 := []int{1}
	nums3 := []int{0, 0, 0, 0, 0}

	t.Run("Testing normal case", func(t *testing.T) {
		want := 5

		got := removeDuplicates(nums1)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Testing single element array", func(t *testing.T) {
		want := 1

		got := removeDuplicates(nums2)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Testing all duplicates", func(t *testing.T) {
		want := 1

		got := removeDuplicates(nums3)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
