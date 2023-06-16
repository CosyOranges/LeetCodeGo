package arrays

import "testing"

func TestIsPalindrome(t *testing.T) {
	empty := " "
	punctuation := "A man, a plan, a canal: Panama"

	t.Run("Test empty string", func(t *testing.T) {
		want := true

		got := isPalindrome(empty)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test complex string", func(t *testing.T) {
		want := true

		got := isPalindrome(punctuation)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
