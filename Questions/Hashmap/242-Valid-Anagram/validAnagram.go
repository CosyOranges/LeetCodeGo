package hashmap

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[rune]int)

	for _, x := range s {
		sMap[x]++
	}

	for _, x := range t {
		if _, ok := sMap[x]; ok {
			if sMap[x] > 0 {
				sMap[x]--
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func main() {
	s := "ab"
	t := "a"

	fmt.Printf("Ans: %v\n", isAnagram(s, t))
}
