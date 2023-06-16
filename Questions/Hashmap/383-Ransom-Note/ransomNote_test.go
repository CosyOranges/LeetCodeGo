package arrays

import "testing"

func TestRansomNote(t *testing.T) {
	magazine := "Use this string to make a ransome note"

	t.Run("Test", func(t *testing.T) {
		want := "thin rake"

		if !ransomNote(want, magazine) {
			t.Errorf("Could not make %v.", want)
		}
	})
}
