package arrays

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	test := []int{-1, 0, 1, 2, -1, -4}

	t.Run("Test standard array with duplicates", func(t *testing.T) {
		want := [][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		}

		got := threesum(test)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
