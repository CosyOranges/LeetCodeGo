package arrays

import "testing"

func TestJump(t *testing.T) {
	impossible := []int{3, 2, 1, 0, 4}
	possible := []int{3, 2, 2, 1, 4}
	short := []int{0}

	t.Run("Testing Possible", func(t *testing.T) {
		got := canJump(possible)

		if !got {
			t.Errorf("Got false, expected true")
		}
	})

	t.Run("Testing Impossible", func(t *testing.T) {
		got := canJump(impossible)

		if got {
			t.Errorf("Got true, expected false")
		}
	})

	t.Run("Testing short case", func(t *testing.T) {
		got := canJump(short)

		if !got {
			t.Errorf("Got false, expected true")
		}
	})
}
