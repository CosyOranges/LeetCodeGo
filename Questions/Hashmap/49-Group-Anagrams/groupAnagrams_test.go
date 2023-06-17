package hashmap

import (
	"testing"
)

func TestGroupAnagram(t *testing.T) {
	test := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	t.Run("Test", func(t *testing.T) {
		want := [][]string{
			{"eat", "tea", "ate"},
			{"tan", "nat"},
			{"bat"},
		}

		got := groupAnagrams(test)

		if len(want) != len(got) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func TestGroupAnagramII(t *testing.T) {
	test := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	t.Run("Test", func(t *testing.T) {
		want := [][]string{
			{"eat", "tea", "ate"},
			{"tan", "nat"},
			{"bat"},
		}

		got := groupAnagramsII(test)

		if len(want) != len(got) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
