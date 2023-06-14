package arrays

import "testing"

func TestReverseWords(t *testing.T) {
	empty := ""
	oneWord := "test"
	singleLetterEdgeCase := "a good   example"

	t.Run("Test empty string", func(t *testing.T) {
		want := ""

		got := ReverseWords(empty)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test one word string", func(t *testing.T) {
		want := "test"

		got := ReverseWords(oneWord)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test single letter edge case string", func(t *testing.T) {
		want := "example good a"

		got := ReverseWords(singleLetterEdgeCase)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
