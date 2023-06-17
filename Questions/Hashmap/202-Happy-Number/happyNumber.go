package hashmap

import "fmt"

func isHappy(num int) bool {
	digits := getDigits(num, []int{})
	intMap := make(map[int]int)
	ans := true
	for ans {
		sum := 0
		for _, x := range digits {
			sum += (x * x)
		}
		intMap[sum]++
		if intMap[sum] > 1 {
			ans = false
		} else if sum == 1 {
			break
		}
		digits = getDigits(sum, []int{})
	}
	return ans
}

func getDigits(x int, digits []int) []int {
	for x != 0 {
		i := x % 10
		digits = append([]int{i}, digits...)
		return getDigits(x/10, digits)
	}
	return digits
}

func main() {
	fmt.Printf("Ans: %v\n", isHappy(2))
}
