package arrays

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 2, 7, 7, 11, 15}
	empty := []int{}

	t.Run("Test with empty array", func(t *testing.T) {
		want := []int{0, 0}

		got := twoSum(empty, 0)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test with normal array", func(t *testing.T) {
		want := []int{3, 1}

		got := twoSum(nums, 9)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test with normal array no solution", func(t *testing.T) {
		want := []int{0, 0}

		got := twoSum(nums, -1)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
