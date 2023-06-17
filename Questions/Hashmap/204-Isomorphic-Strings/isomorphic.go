package arrays

import "fmt"

func isomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap, tMap := make(map[byte]byte), make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		_, oks := sMap[s[i]]
		_, okt := tMap[t[i]]
		if !oks {
			sMap[s[i]] = t[i]
		}
		if !okt {
			tMap[t[i]] = s[i]
		}

		if sMap[s[i]] != t[i] || tMap[t[i]] != s[i] {
			fmt.Printf("here s: %c | t: %c | index: %v\n", s[i], t[i], i)
			fmt.Printf("MapSVal: %c | t: %c\n", sMap[s[i]], t[i])
			fmt.Printf("MapTVal: %c | s: %c\n", tMap[t[i]], s[i])
			return false
		}
	}

	return true
}

func main() {
	s := "paper"
	t := "title"

	fmt.Printf("Isomorphic: %v\n", isomorphic(s, t))
}
