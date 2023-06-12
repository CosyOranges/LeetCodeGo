package arrays

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	negative := []int{-1, 1, 0, -3, 3}

	t.Run("Test standard array", func(t *testing.T) {
		want := []int{24, 12, 8, 6}

		got := productExceptSelf(nums)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test negative array", func(t *testing.T) {
		want := []int{0, 0, 9, 0, 0}

		got := productExceptSelf(negative)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
