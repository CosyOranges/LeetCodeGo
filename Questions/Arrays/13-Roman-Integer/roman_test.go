package arrays

import "testing"

func TestRomanToInt(t *testing.T) {
	first := "III"
	second := "MCMXCIV"

	t.Run("Test simple incrementing numeral", func(t *testing.T) {
		want := 3

		got := romanToInt(first)

		if got != want {
			t.Errorf("got: %v, want %v", got, want)
		}
	})

	t.Run("Test complex numeral", func(t *testing.T) {
		want := 1994

		got := romanToInt(second)

		if got != want {
			t.Errorf("got: %v, want %v", got, want)
		}
	})
}

func TestSwitchRomanCase(t *testing.T) {
	t.Run("Test with valid Letter", func(t *testing.T) {
		got := romanSwitchCase('M')

		if got != 1000 {
			t.Errorf("got: %v, want: %v", got, 1000)
		}
	})

	t.Run("Test with invalid Letter", func(t *testing.T) {
		got := romanSwitchCase('Z')

		if got != 0 {
			t.Errorf("got: %v, want: %v", got, 0)
		}
	})
}
