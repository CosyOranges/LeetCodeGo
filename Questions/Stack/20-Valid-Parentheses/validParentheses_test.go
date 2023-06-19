package stack

import "testing"

func TestValidParentheses(t *testing.T) {
	bracketsValid := "[][][]"
	bracketsInValid := "[)[][]"

	t.Run("Test valid", func(t *testing.T) {
		if !validParentheses(bracketsValid) {
			t.Errorf("got: false, want: true")
		}
	})

	t.Run("Test invalid", func(t *testing.T) {
		if validParentheses(bracketsInValid) {
			t.Errorf("got: true, want: false")
		}
	})
}
