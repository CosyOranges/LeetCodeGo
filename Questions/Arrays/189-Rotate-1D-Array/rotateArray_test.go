package arrays

import (
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	nums1 := []int{1, 2, 3, 4, 5, 6}

	t.Run("Testing normal case", func(t *testing.T) {
		want := []int{5, 6, 1, 2, 3, 4}

		rotate(nums1, 2)

		if !reflect.DeepEqual(want, nums1) {
			t.Errorf("got: %v, want: %v", nums1, want)
		}
	})

	t.Run("Testing k > len(nums) case", func(t *testing.T) {
		// Increment by a further 2 because we are not deep copying nums and
		// the previous test will have already rotated the array
		want := []int{3, 4, 5, 6, 1, 2}

		rotate(nums1, 2)

		if !reflect.DeepEqual(want, nums1) {
			t.Errorf("got: %v, want: %v", nums1, want)
		}
	})
}

func TestReverse(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}

	t.Run("Testing reverse full array", func(t *testing.T) {
		want := []int{6, 5, 4, 3, 2, 1}

		reverse(nums, 0, len(nums)-1)

		if !reflect.DeepEqual(want, nums) {
			t.Errorf("got: %v, want: %v", nums, want)
		}
	})

	t.Run("Testing reverse slice of array", func(t *testing.T) {
		want := []int{6, 5, 4, 1, 2, 3}

		reverse(nums, 3, len(nums)-1)

		if !reflect.DeepEqual(want, nums) {
			t.Errorf("got: %v, want: %v", nums, want)
		}
	})
}
