package hashmap

import "testing"

func TestIsAnagram(t *testing.T) {
	s1 := "ab"
	t1 := "a"
	s2 := "anagram"
	t2 := "margana"

	t.Run("Test invalid case", func(t *testing.T) {
		if isAnagram(s1, t1) {
			t.Errorf("got: true, want: false")
		}
	})

	t.Run("Test valid case", func(t *testing.T) {
		if !isAnagram(s2, t2) {
			t.Errorf("got: false, want: true")
		}
	})
}
