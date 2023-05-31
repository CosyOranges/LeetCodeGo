package arrays

import "testing"

func TestRemoveDuplicatesII(t *testing.T) {
	// Set up some test cases
	nums1 := []int{1, 1, 1, 2, 2, 3}
	nums2 := []int{1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 4, 5}
	nums3 := []int{1, 1}

	// Set tests to be run i.e. t.Run("description", func(t *testing.T) {...})
	t.Run("Test standard arrays", func(t *testing.T) {
		want := 5

		got := removeDuplicatesII(nums1)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}

		want = 6

		got = removeDuplicatesII(nums2)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test short arrays", func(t *testing.T) {
		want := 2

		got := removeDuplicatesII(nums3)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
