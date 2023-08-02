package backtracking

import (
	"reflect"
	"testing"
)

func TestLetterBFS(t *testing.T) {
	t.Run("Test all 1s", func(t *testing.T) {
		want := []string{}

		got := lettersBFS("11111")

		if got != nil {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test 123", func(t *testing.T) {
		want := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}

		got := lettersBFS("123")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
