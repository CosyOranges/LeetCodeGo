package arrays

import "testing"

func TestLongestConsecutiveSequence(t *testing.T) {
	nums := []int{1, 2, 0, 1}
	t.Run("Test Map version with duplicates", func(t *testing.T) {
		want := 3

		got := longestConsecutiveSequenceMap(nums)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test non-map version with duplicates", func(t *testing.T) {
		want := 3

		got := longestConsecutiveSequence(nums)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
