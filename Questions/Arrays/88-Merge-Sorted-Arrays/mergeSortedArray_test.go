package arrays

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	nums3 := []int{1, 2, 3}
	nums4 := []int{4, 0, 0, 0, 0, 0}
	nums5 := []int{1, 2, 3, 5, 6}
	empty := []int{0, 0, 0}
	t.Run("Test two sorted arrays", func(t *testing.T) {
		want := []int{1, 2, 2, 3, 5, 6}

		got := mergeArrays(nums1, 3, nums2, 3)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test with an empty first array", func(t *testing.T) {
		want := []int{2, 5, 6}

		got := mergeArrays(empty, 0, nums2, 3)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test with an empty second array", func(t *testing.T) {
		want := []int{1, 2, 3}

		got := mergeArrays(nums3, 3, empty, 0)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Final test", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5, 6}

		got := mergeArrays(nums4, 1, nums5, 5)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
