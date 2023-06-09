package arrays

import "testing"

func TestJump2(t *testing.T) {
	nums1 := []int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}

	t.Run("Test significant array", func(t *testing.T) {
		want := 2

		got := jump(nums1)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
