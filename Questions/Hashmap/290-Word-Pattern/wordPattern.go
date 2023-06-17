package hashmap

import (
	"fmt"
	"reflect"
)

func wordPattern(pattern string, s string) bool {
	pMap, sMap := make(map[byte][]int), make(map[string][]int)

	wordNumb, prev := 0, 0

	for i := 0; i < len(s); i++ {
		t := wordNumb % len(pattern)
		if s[i] == ' ' || i == len(s)-1 {
			pMap[pattern[t]] = append(pMap[pattern[t]], wordNumb)
			if i == len(s)-1 {
				sMap[s[prev:]] = append(sMap[s[prev:]], t)
				if !reflect.DeepEqual(sMap[s[prev:]], pMap[pattern[len(pattern)-1]]) {
					return false
				}
			} else {
				sMap[s[prev:i]] = append(sMap[s[prev:i]], t)
				if !reflect.DeepEqual(sMap[s[prev:i]], pMap[pattern[t]]) {
					return false
				}
				prev = i + 1
				wordNumb++

			}
		}
	}
	fmt.Printf("Pattern Map: %v\n", pMap)
	fmt.Printf("String Map: %v\n", sMap)

	return true
}

func main() {
	pat := "aaa"
	s := "aa aa aa aa"

	fmt.Printf("Ans: %v\n", wordPattern(pat, s))
}
