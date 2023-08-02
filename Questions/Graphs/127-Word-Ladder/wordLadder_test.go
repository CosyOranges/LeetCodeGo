package graphs

import "testing"

func TestWordLadder(t *testing.T) {
	t.Run("Test with a valid ladder", func(t *testing.T) {
		want := 5

		got := wordLadder("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test with an invalid ladder", func(t *testing.T) {
		want := 0

		got := wordLadder("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"})

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
