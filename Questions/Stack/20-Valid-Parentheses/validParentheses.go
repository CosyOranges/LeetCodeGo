package stack

func validParentheses(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	pairs := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	stack := []rune{}

	for _, r := range s {
		if _, ok := pairs[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != r {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}

// func main() {
// 	brackets := "{}{]{}{}"

// 	fmt.Printf("Ans: %v\n", validParentheses(brackets))
// }
