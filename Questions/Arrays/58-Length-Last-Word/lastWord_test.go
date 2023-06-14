package arrays

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	t.Run("Test empty string", func(t *testing.T) {
		got := LengthOfLastWord("")

		if got != 0 {
			t.Errorf("got: %v, want: 0", got)
		}
	})

	t.Run("Test valid string", func(t *testing.T) {
		got := LengthOfLastWord("  testtwo test    ")

		if got != 4 {
			t.Errorf("got: %v, want: 4", got)
		}
	})
}
