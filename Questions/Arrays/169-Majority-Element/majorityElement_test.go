package arrays

import "testing"

func TestMajorityElements(t *testing.T) {
	nums1 := []int{2, 2, 1, 1, 1, 2, 2}

	t.Run("Test normal case", func(t *testing.T) {
		want := 2

		got := majorityElement(nums1)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
