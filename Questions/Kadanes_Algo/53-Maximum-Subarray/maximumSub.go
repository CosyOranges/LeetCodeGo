package kadanes

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	localSum, globalSum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > localSum+nums[i] {
			localSum = nums[i]
		} else {
			localSum += nums[i]
		}

		if globalSum < localSum {
			globalSum = localSum
		}
	}

	return globalSum
}

// func main() {
// 	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

// 	fmt.Printf("Ans: %v\n", maxSubArray(nums))
// }
