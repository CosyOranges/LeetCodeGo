package arrays

import (
	"fmt"
)

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	longestPrefix := ""

	// To solve this we can continuously loop through the array on every loop
	// we compare the same index of each word using strs[0] as our source of truth
	// if we every encounter a character that doesn't match up with strs[0][i] then
	// return the current longest prefix
	// Or if i is equal to the length of a word then also return current longest prefix
	// because the longest prefix cannot be longer than any word
	i := 0
	loop := true
	for loop {
		for _, s := range strs {
			if i == len(s) || strs[0][i] != s[i] {
				return longestPrefix
			}
		}
		longestPrefix += string(strs[0][i])
		i++
	}

	return longestPrefix
}

func main() {
	test := []string{"flower", "flow", "flight"}
	fmt.Printf("Ans: %v\n", LongestCommonPrefix(test))
}
