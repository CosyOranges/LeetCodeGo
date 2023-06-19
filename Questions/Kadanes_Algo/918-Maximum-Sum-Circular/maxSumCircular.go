package kadanes

import (
	"fmt"
)

// Kadane's Normal Sub Array
// func maxSubarraySum(nums []int) int {
// 	globalMaxSum := math.MinInt
// 	localMaxSum := math.MinInt

// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] > localMaxSum+nums[i] {
// 			localMaxSum = nums[i]
// 		} else {
// 			localMaxSum += nums[i]
// 		}

// 		if localMaxSum > globalMaxSum {
// 			globalMaxSum = localMaxSum
// 		}
// 	}

// 	return globalMaxSum
// }

// Kadane's Circular Sub Array
func maxSubarraySumCircular(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	// find the total value of the array
	total := 0
	for _, x := range nums {
		total += x
	}

	// To simulate a circular array the maximum contiguous array could be (denoted by x's):
	// [x, x, x, _, _, _, x, x]
	// From which it can be reasoned that the sub-array of [_, _, _] is the minimum
	// Thus to find the maximum sub_array we need to also use Kidanes to find the minimum sub_array
	globalMaxSum, localMaxSum, globalMinSum, localMinSum := nums[0], nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		// Find the max
		if nums[i] > localMaxSum+nums[i] {
			localMaxSum = nums[i]
		} else {
			localMaxSum += nums[i]
		}
		if localMaxSum > globalMaxSum {
			globalMaxSum = localMaxSum
		}

		// Find the min
		if nums[i] < localMinSum+nums[i] {
			localMinSum = nums[i]
		} else {
			localMinSum += nums[i]
		}
		if localMinSum < globalMinSum {
			globalMinSum = localMinSum
		}

	}
	if globalMinSum == total {
		return globalMaxSum
	}

	if total-globalMinSum > globalMaxSum {
		return total - globalMinSum
	}
	return globalMaxSum
}

func main() {
	nums := []int{1, -2, 3, -2}
	fmt.Printf("Ans: %v\n", maxSubarraySumCircular(nums))
}
