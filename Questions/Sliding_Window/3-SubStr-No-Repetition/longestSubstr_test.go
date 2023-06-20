package sliding_window

import (
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	test1 := "abcabcdefbb"
	test2 := "a"
	test3 := ""

	t.Run("Test normal string", func(t *testing.T) {
		want := 6

		got := longestSubstring(test1)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test short string", func(t *testing.T) {
		want := 1

		got := longestSubstring(test2)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("Test no string", func(t *testing.T) {
		want := 0

		got := longestSubstring(test3)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
