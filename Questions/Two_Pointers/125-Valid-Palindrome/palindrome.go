package arrays

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, "")
	s = strings.ToLower(s)
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Printf("Is Palindrome: %v\n", isPalindrome("A man, a plan, a canal: Panama"))
}
