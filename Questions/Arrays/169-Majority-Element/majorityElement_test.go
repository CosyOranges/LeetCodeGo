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

func TestMajorityElementConstantSpace(t *testing.T) {
	nums1 := []int{2, 2, 1, 1, 1, 2, 2}
	nums2 := []int{1, 1, 2, 2, 3, 3}

	t.Run("Test normal case", func(t *testing.T) {
		want := 2

		got := majorityElementConstantSpace(nums1)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test case with equal number", func(t *testing.T) {
		want := 3

		// Should return the final element
		got := majorityElementConstantSpace(nums2)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
