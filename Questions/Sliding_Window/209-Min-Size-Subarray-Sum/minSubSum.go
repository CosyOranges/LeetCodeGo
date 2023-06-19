package sliding_window

import "fmt"

// Brute Force Approach
// func minSubArray(nums []int, target int) int {
// 	sum, window := 0, 1

// 	for window < len(nums)+1 {
// 		start, end := 0, window

// 		fmt.Printf("Window: %v | start: %v | end: %v \n", window, start, end)
// 		for end < len(nums)+1 {
// 			for i := start; i < end; i++ {
// 				sum += nums[i]
// 			}
// 			if sum >= target {
// 				return window
// 			}
// 			sum = 0
// 			end++
// 			start++
// 		}
// 		window++
// 	}

// 	return 0
// }

func minSubArray(nums []int, target int) int {
	l, sum, minLength := 0, 0, len(nums)+1

	for r := 0; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target {
			if minLength > (r-l)+1 {
				minLength = (r - l) + 1
			}
			sum -= nums[l]
			l++
		}
	}

	// We did not find a subarray that satisfied the target
	if minLength == len(nums)+1 {
		return 0
	}

	return minLength
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}

	fmt.Printf("Ans: %v\n", minSubArray(nums, 7))
}
