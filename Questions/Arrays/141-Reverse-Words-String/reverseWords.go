package arrays

import "fmt"

func ReverseWords(words string) string {
	ans := ""
	mark := -1
	for i := len(words) - 1; i >= 0; i-- {
		if words[i] != ' ' && mark == -1 {
			mark = i
		}
		if (words[i] == ' ' && mark != -1) || (i == 0 && mark != -1) {
			temp := words[i : mark+1]
			if ans == "" {
				if temp[0] == ' ' {
					ans += temp[1:]
				} else {
					ans += temp
				}
			} else {
				if temp[0] == ' ' {
					ans += " " + temp[1:]
				} else {
					ans += " " + temp
				}
			}

			mark = -1
		}
	}

	return ans
}

func main() {
	fmt.Printf("Reverse: %v STOP\n", ReverseWords("a good    example"))
}
