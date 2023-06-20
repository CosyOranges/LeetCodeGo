package sliding_window

import "fmt"

func longestSubstring(s string) int {
	left, right := 0, 0
	sMap := make(map[byte]int)
	sublength := 0

	for right < len(s) {
		if _, ok := sMap[s[right]]; ok {
			if right-left > sublength {
				sublength = right - left
			}
			delete(sMap, s[left])
			left++
		} else {
			sMap[s[right]]++
			right++
		}
	}

	if len(sMap) > sublength {
		return len(sMap)
	}
	return sublength
}

/*
"abcabcdefbb"
*/
func main() {
	fmt.Printf("Ans: %v\n", longestSubstring("abcabcdefbb"))
}
