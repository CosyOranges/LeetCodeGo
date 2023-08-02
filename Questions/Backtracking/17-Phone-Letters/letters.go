package backtracking

import (
	"fmt"
	"strings"
)

var phonePad = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digit string) []string {
	var ans []string = []string{}
	if len(digit) == 0 {
		return nil
	}

	/*
		digit: 23
		- each combo can only be len(digit)
			- but for each 1 contained in the digit, the length of each combo is reduced by one
			- therefore goal length =  len(digit) - count_of_ones

		- we want to set up a series of strings
			"a" "b" "c"
		- We have handled the first digit and so can remove it from the string
			- There fore we are left with digit[1:] or just 3
		- Next we append all combos from 3 to each combo from 2
			- creating: "ad" "ae" "af" "bd" .... etc

	*/
	strings.ReplaceAll(digit, "1", "")

	fmt.Printf("Digit: %v\n", digit)
	ans = builder(digit)

	return ans
}

func builder(digit string) []string {
	combo := []string{}

	if len(digit) == 0 {
		return nil
	}

	first := string(digit[0])

	if len(digit) == 1 {
		return phonePad[first]
	}

	for _, l := range phonePad[first] {
		for _, rest := range builder(digit[1:]) {
			combo = append(combo, l+rest)
		}
	}

	return combo
}

func lettersBFS(digits string) []string {
	ans := []string{}
	digits = strings.Replace(digits, "1", "", -1)
	if len(digits) == 0 {
		return nil
	}

	// Remove all 1 digits

	// Add first set of digits
	ans = append(ans, phonePad[string(digits[0])]...)

	// BFS on the digits
	for i := 1; i < len(digits); i++ {
		letters := phonePad[string(digits[i])]

		n := len(ans)
		// Now loop through each letter and append it plus the current letters to the end
		for j := 0; j < n; j++ {
			for _, letter := range letters {
				ans = append(ans, ans[j]+letter)
			}
		}

		// Remove the n front digits
		ans = ans[n:]
	}

	return ans
}

func main() {
	fmt.Printf("Combos: %v\n", letterCombinations("23"))
	fmt.Printf("Combos: %v\n", lettersBFS("123"))
}
