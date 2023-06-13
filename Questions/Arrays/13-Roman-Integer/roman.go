package arrays

import "fmt"

func romanToInt(s string) int {
	ans := 0

	for i := 0; i < len(s); i++ {
		if i+1 <= len(s)-1 {
			if s[i] == 'I' && (s[i+1] == 'V' || s[i+1] == 'X') {
				// Then we have I infront of either V or X
				fmt.Printf("'I' infront of %v\n", s[i])
				switch s[i+1] {
				case 'V':
					ans += 4
				case 'X':
					ans += 9
				}
				i++
			} else if s[i] == 'X' && (s[i+1] == 'L' || s[i+1] == 'C') {
				fmt.Printf("'X' infront of %v\n", s[i])
				switch s[i+1] {
				case 'L':
					ans += 40
				case 'C':
					ans += 90
				}
				i++
			} else if s[i] == 'C' && (s[i+1] == 'D' || s[i+1] == 'M') {
				fmt.Printf("'C' infront of %v\n", s[i])
				switch s[i+1] {
				case 'D':
					ans += 400
				case 'M':
					ans += 900
				}
				i++
			} else {
				// Just a regular add
				ans += romanSwitchCase(s[i])
			}
		} else {
			ans += romanSwitchCase(s[i])
		}
		fmt.Printf("Num so far: %v\n", ans)
	}
	return ans
}

func romanSwitchCase(s byte) int {
	switch s {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}

func main() {
	roman := "MCMXCIV"

	fmt.Printf("ANS: %v\n", romanToInt(roman))
}
