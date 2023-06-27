package hashmap

import "fmt"

func parenthesisSet(n int) []string {

	// I'm going to handle 0 and 1 as a base case for the initial function
	// For the helper function we will only care about 0
	if n == 0 {
		return []string{""}
	} else if n == 1 {
		return []string{"()"}
	}

	// Initialise our solution array
	ans := []string{}

	// Create a combos map, a Map is used here instead of a set because Go does not have a native set
	combos := map[string]int{}

	// We will always have a starting point of "()"
	temp := "()"

	// No return from the helper function because Go passes by reference and so any changes
	// made to combos in the function will be available to us here without an explicit return
	// I'm also going to carry the original n to use to check if the combo created is of the right length
	// This could be done a different way by checking after the combo is created whether n-1 == 0
	helper(combos, n, n-1, temp)

	// Grab all the combos from the map
	for key := range combos {
		ans = append(ans, key)
	}

	return ans
}

func helper(combos map[string]int, original int, n int, temp string) {
	// If n is 0 return nothing left to be done
	if n == 0 {
		return
	}
	// Check for an append of '()' permutation to the string
	if _, ok := combos[temp+"()"]; !ok {
		// We only want to add to the map if it has enough characters in it
		if len(temp+"()") == 2*original {
			combos[temp+"()"]++
		}
	}
	// Recursively explore appending '()' to the string
	helper(combos, original, n-1, temp+"()")

	// Check for a prepend of '()' permutation to the string
	if _, ok := combos["()"+temp]; !ok {
		if len("()"+temp) == 2*original {
			combos["()"+temp]++
		}
	}
	// Recursively explore prepending '()'
	helper(combos, original, n-1, "()"+temp)

	// Check for a prepend of '(' and an append of ')'
	if _, ok := combos["("+temp+")"]; !ok {
		if len("("+temp+")") == 2*original {
			combos["("+temp+")"]++
		}
	}
	// Recursively explore prepending and appending
	helper(combos, original, n-1, "("+temp+")")

	// Explicitly return once no more permutations are available to us
	return
}

func main() {
	fmt.Printf("Ans: %v\n", parenthesisSet(3))
}
