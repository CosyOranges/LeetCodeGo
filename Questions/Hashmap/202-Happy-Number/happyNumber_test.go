package hashmap

import "testing"

func TestIsHappy(t *testing.T) {
	t.Run("Test unhappy number", func(t *testing.T) {
		if isHappy(2) {
			t.Errorf("got: true, want: false")
		}
	})

	t.Run("Test happy number", func(t *testing.T) {
		if !isHappy(19) {
			t.Errorf("got: want, want: true")
		}
	})
}
