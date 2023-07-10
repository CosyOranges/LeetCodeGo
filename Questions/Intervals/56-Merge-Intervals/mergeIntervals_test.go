package intervals

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	test1 := [][]int{{1, 4}, {2, 3}, {8, 10}, {9, 20}}
	test2 := [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}

	t.Run("Test Case 1", func(t *testing.T) {
		want := [][]int{{1, 4}, {8, 20}}

		got := mergeIntervals(test1)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test Case 2", func(t *testing.T) {
		want := [][]int{{1, 10}}

		got := mergeIntervals(test2)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
