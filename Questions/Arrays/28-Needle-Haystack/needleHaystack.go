package arrays

import "fmt"

func strStr(haystack string, needle string) int {
	n := len(needle)
	i := 0

	for n < len(haystack)+1 {
		if haystack[i:n] == needle {
			return i
		}
		n++
		i++
	}

	return -1
}

func main() {
	fmt.Printf("Ans: %v\n", strStr("sadbutsad", "sad"))
}
