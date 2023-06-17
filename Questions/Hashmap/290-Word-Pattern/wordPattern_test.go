package hashmap

import "testing"

func TestWordPattern(t *testing.T) {
	p1 := "abba"
	s1 := "dog fish fish dog"

	p2 := "aaa"
	s2 := "aa aa aa aa"
	s3 := "aa aa aa aa aa aa"

	t.Run("Test valid pattern string pair", func(t *testing.T) {
		if !wordPattern(p1, s1) {
			t.Errorf("got: false, want: true")
		}
	})

	t.Run("Test invalid pattern string pair", func(t *testing.T) {
		if wordPattern(p2, s2) {
			t.Errorf("got: true, want: false")
		}
	})

	t.Run("Test repeating pattern string pair", func(t *testing.T) {
		if wordPattern(p2, s3) {
			t.Errorf("got: false, want: true")
		}
	})
}
