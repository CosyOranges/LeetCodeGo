package arrays

import (
	"reflect"
	"testing"
)

func TestFullJustify(t *testing.T) {
	test := []string{"Listen", "to", "many,", "speak", "to", "a", "few."}
	empt := []string{""}

	t.Run("Test string with word as long as maxWidth", func(t *testing.T) {
		want := []string{
			"Listen",
			"to    ",
			"many, ",
			"speak ",
			"to   a",
			"few.  ",
		}

		got := FullJustify(test, 6)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test empty string", func(t *testing.T) {
		want := []string{
			" ",
		}

		got := FullJustify(empt, 1)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func TestLineBuilder(t *testing.T) {
	one := "test"
	two := "test test"

	t.Run("Test as last line", func(t *testing.T) {
		want := "test    "

		got := lineBuilder(3, one+" ", true)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test as normal line", func(t *testing.T) {
		want := "test    test"

		got := lineBuilder(3, two, false)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
