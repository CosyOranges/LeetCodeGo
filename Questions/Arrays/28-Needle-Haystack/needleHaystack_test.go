package arrays

import "testing"

func TestStrStr(t *testing.T) {
	s := "sadbutsad"

	t.Run("Test no occurence", func(t *testing.T) {
		want := -1

		got := strStr(s, "xxx")

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test occurence", func(t *testing.T) {
		want := 0

		got := strStr(s, "sad")

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
