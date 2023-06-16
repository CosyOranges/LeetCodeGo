package arrays

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	test := []int{-10, -8, -2, 1, 2, 5, 6}

	t.Run("Test negative test case", func(t *testing.T) {
		want := []int{3, 5}

		got := twoSum(test, 0)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
