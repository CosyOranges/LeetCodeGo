package arrays

import "fmt"

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	charsFoundLeft := 0
	charsFoundRight := len(s) - 1

	left := 0
	right := len(t) - 1

	for left <= right {
		if t[left] == s[charsFoundLeft] {
			fmt.Printf("Found: %c, pos: %v\n", t[left], left)
			charsFoundLeft++
			if charsFoundLeft > charsFoundRight {
				return true
			}
		}

		// We add && left != right, because we dont want to double count a char from t
		if t[right] == s[charsFoundRight] && left != right {
			fmt.Printf("Found: %c, pos: %v\n", t[right], right)
			charsFoundRight--
			if charsFoundLeft > charsFoundRight {
				return true
			}
		}
		left++
		right--
	}
	return false
}

func main() {
	fmt.Printf("Substring: %v\n", isSubsequence("b", "ahbdc"))
}
