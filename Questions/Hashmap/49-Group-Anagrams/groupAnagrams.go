package hashmap

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	sMap := make(map[string][]string)
	for _, x := range strs {
		temp := sortString(x)
		sMap[temp] = append(sMap[temp], x)
	}

	for _, value := range sMap {
		ans = append(ans, value)
	}
	return ans
}

type sortRunes []rune

func (s sortRunes) Less(i int, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)

	sort.Sort(sortRunes(r))
	return string(r)
}

// Another Solution is to create a bespoke hashing solution
type Key [26]byte

func hash(str string) Key {
	key := Key{}

	for i := range str {
		key[str[i]-'a']++
	}

	return key
}

func groupAnagramsII(str []string) [][]string {
	ans := [][]string{}
	sMap := make(map[Key][]string)

	for _, x := range str {
		sMap[hash(x)] = append(sMap[hash(x)], x)
	}

	for _, value := range sMap {
		ans = append(ans, value)
	}

	return ans
}

func main() {
	test := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Printf("Ans: %v\n", groupAnagramsII(test))
}
