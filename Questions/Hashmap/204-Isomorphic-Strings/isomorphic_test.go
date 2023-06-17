package hashmap

import "testing"

func TestIsomorphic(t *testing.T) {
	s1 := "paper"
	s2 := "title"
	s3 := "add"
	s4 := "abc"

	t.Run("Test isomorphic strings", func(t *testing.T) {
		if !isomorphic(s1, s2) {
			t.Errorf("got: false, want: true")
		}
	})

	t.Run("Test non-isomorphic strings", func(t *testing.T) {
		if isomorphic(s3, s4) {
			t.Errorf("got: true, want: false")
		}
	})
}
