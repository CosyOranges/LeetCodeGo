package arrays

import "fmt"

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))

	// To solve this we can make use of two arrays that slide from either side
	i, j := 1, len(nums)-2
	left, right := make([]int, len(nums)), make([]int, len(nums))
	left[0], right[j+1] = 1, 1
	for i < len(nums) && j >= 0 {
		left[i] = nums[i-1] * left[i-1]
		right[j] = nums[j+1] * right[j+1]
		i++
		j--
	}

	fmt.Printf("Left: %v | Right: %v\n", left, right)
	for i = 0; i < len(nums); i++ {
		ans[i] = left[i] * right[i]
	}

	return ans
}

func main() {
	nums := []int{1, 2, 3, 4}

	fmt.Printf("Ans: %v\n", productExceptSelf(nums))
}
