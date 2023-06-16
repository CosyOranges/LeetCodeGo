package arrays

import "testing"

func TestIsSubsequence(t *testing.T) {
	base1 := "sdfsdfa"
	base2 := "ahbdc"
	base3 := ""

	t.Run("Test with empty substring", func(t *testing.T) {
		if !isSubsequence("", base1) {
			t.Error("got: false, want: true")
		}
	})

	t.Run("Test with invalid substring", func(t *testing.T) {
		if isSubsequence("abbc", base2) {
			t.Error("got: true, want: false")
		}
	})

	t.Run("Test with empty string", func(t *testing.T) {
		if isSubsequence("abbc", base3) {
			t.Error("got: true, want: false")
		}
	})

	t.Run("Test with valid substring", func(t *testing.T) {
		if !isSubsequence("abc", base2) {
			t.Error("got: false, want: true")
		}
	})
}
