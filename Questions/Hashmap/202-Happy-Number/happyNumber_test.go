package hashmap

import (
	"reflect"
	"testing"
)

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

func TestGetDigits(t *testing.T) {
	t.Run("Test single digit", func(t *testing.T) {
		got := getDigits(2, []int{})

		if !reflect.DeepEqual(got, []int{2}) {
			t.Errorf("got: %v, want: %v", got, []int{2})
		}
	})

	t.Run("Test multiple digit", func(t *testing.T) {
		got := getDigits(2222, []int{})

		if !reflect.DeepEqual(got, []int{2, 2, 2, 2}) {
			t.Errorf("got: %v, want: %v", got, []int{2, 2, 2, 2})
		}
	})
}
