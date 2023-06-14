package arrays

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	empty := []string{""}
	noPrefix := []string{"", "b"}
	prefixes := []string{"flower", "flow", "flight"}

	t.Run("Test empty array", func(t *testing.T) {
		want := ""

		got := LongestCommonPrefix(empty)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test no prefix array", func(t *testing.T) {
		want := ""

		got := LongestCommonPrefix(noPrefix)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test prefix array", func(t *testing.T) {
		want := "fl"

		got := LongestCommonPrefix(prefixes)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
