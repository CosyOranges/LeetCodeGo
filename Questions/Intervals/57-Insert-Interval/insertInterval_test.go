package intervals

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	test := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}

	t.Run("Test overlapping on one side of interval", func(t *testing.T) {
		want := [][]int{{1, 2}, {3, 10}, {12, 16}}

		got := insert(test, []int{4, 8})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test overlap on both sides", func(t *testing.T) {
		want := [][]int{{0, 10}, {12, 16}}

		got := insert(test, []int{0, 8})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
