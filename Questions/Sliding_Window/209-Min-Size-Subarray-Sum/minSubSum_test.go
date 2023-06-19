package sliding_window

import "testing"

func TestMinSubArray(t *testing.T) {
	standard := []int{2, 3, 1, 2, 4, 3}
	short := []int{4, 1, 4}
	invalid := []int{1, 1, 1, 1, 1}

	t.Run("Test standard", func(t *testing.T) {
		want := 2

		got := minSubArray(standard, 7)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test short", func(t *testing.T) {
		want := 1

		got := minSubArray(short, 4)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test invalid", func(t *testing.T) {
		want := 0

		got := minSubArray(invalid, 11)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
