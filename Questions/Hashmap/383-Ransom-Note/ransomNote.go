package hashmap

import "fmt"

func ransomNote(ransomNote string, magazine string) bool {
	ans := true
	sMap := make(map[rune]int)

	for _, x := range magazine {
		sMap[x]++
	}

	for _, x := range ransomNote {
		if _, ok := sMap[x]; !ok {
			return false
		} else if sMap[x] == 0 {
			return false
		} else {
			sMap[x]--
		}
	}

	return ans
}

func main() {
	magazine := "aab"

	fmt.Printf("Can? %v\n", ransomNote("aaa", magazine))
}
