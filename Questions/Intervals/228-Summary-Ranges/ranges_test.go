package intervals

import (
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	nums := []int{0, 1, 2, 4, 5, 7}

	t.Run("Normal case", func(t *testing.T) {
		want := []string{
			"0->2",
			"4->5",
			"7",
		}

		got := summaryRanges(nums)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Invalid case", func(t *testing.T) {
		want := []string{}

		got := summaryRanges([]int{})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
