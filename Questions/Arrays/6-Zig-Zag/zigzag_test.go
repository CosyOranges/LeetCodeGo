package arrays

import "testing"

func TestZigZag(t *testing.T) {
	s := "TESTINGTHISTEST"

	t.Run("Test with no rows", func(t *testing.T) {
		want := "TESTINGTHISTEST"

		got := ZigZag(s, 0)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test with 4 rows", func(t *testing.T) {
		want := "TGEENTTSSIHSTTI"

		got := ZigZag(s, 4)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
