package hashmap

import (
	"reflect"
	"testing"
)

func TestParenthesisSet(t *testing.T) {
	complex := []string{
		"()()()()",
		"(())()()",
		"()()(())",
		"((()))()",
		"()((()))",
		"(((())))",
		"(()()())",
		"((()()))",
		"((())())",
		"(()(()))",
		"()(())()",
		"(()())()",
		"()(()())",
	}
	t.Run("Test edge cases", func(t *testing.T) {
		want := []string{""}

		got := parenthesisSet(0)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

		got = parenthesisSet(1)
		want = []string{"()"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test input of 4", func(t *testing.T) {
		testMap := map[string]int{}
		for _, x := range complex {
			testMap[x]++
		}

		got := parenthesisSet(4)

		for _, x := range got {
			if _, ok := testMap[x]; !ok {
				t.Errorf("%v not in map", x)
			}
		}
	})
}
