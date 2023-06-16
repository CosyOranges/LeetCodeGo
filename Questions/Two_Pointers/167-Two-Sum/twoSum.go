package arrays

import "fmt"

func twoSum(numbers []int, target int) []int {
	ans := []int{1, 2}
	if len(numbers) <= 2 {
		return ans
	}

	left, right := 0, len(numbers)-1

	for left < right {
		if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		} else if numbers[left]+numbers[right] == target {
			ans[0], ans[1] = left+1, right+1
			return ans
		} else {
			left++
		}
	}

	return ans
}

func main() {
	nums := []int{-10, -8, -2, 1, 2, 5, 6}
	fmt.Printf("Solution: %v\n", twoSum(nums, 0))
}
