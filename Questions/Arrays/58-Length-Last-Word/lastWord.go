package arrays

import "fmt"

func LengthOfLastWord(words string) int {
	ans := 0

	for i := len(words) - 1; i >= 0; i-- {
		if words[i] != ' ' {
			ans++
		} else if words[i] == ' ' && ans != 0 {
			break
		}
	}

	return ans
}

func main() {
	test := "  testtwo  test    "
	fmt.Printf("Length of Last Word: %d\n", LengthOfLastWord(test))
}
