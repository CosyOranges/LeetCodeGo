package arrays

import "testing"

func TestIntToRoman(t *testing.T) {
	t.Run("Test 0", func(t *testing.T) {
		got := intToRoman(0)

		if got != "" {
			t.Errorf("got: %v, want: ''", got)
		}
	})

	t.Run("Test Upper Limit", func(t *testing.T) {
		got := intToRoman(3999)

		if got != "MMMCMXCIX" {
			t.Errorf("got: %v, want: MMMCMXCIX", got)
		}
	})
}
