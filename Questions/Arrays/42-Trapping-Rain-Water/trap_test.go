package arrays

import (
	"testing"
)

func TestTrap(t *testing.T) {
	heights1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	heights2 := []int{0, 1, 2, 3, 4, 5, 6, 7}
	heights3 := []int{0, 0, 0, 0, 0}
	t.Run("Test undulating landscape", func(t *testing.T) {
		want := 6

		got := trap(heights1)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test sloped landscape", func(t *testing.T) {
		want := 0

		got := trap(heights2)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test flat landscape", func(t *testing.T) {
		want := 0

		got := trap(heights3)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
