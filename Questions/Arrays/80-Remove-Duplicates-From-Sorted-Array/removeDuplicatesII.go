package arrays

import "fmt"

func removeDuplicatesII(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	prev1 := nums[0]
	prev2 := nums[1]

	k := 2

	for i := 2; i < len(nums); i++ {
		// fmt.Printf("prev1: %v | prev2: %v | nums[i]: %v | nums[k]: %v\n", prev1, prev2, nums[i], nums[k])
		if prev1 == prev2 && nums[i] != prev2 {
			nums[k] = nums[i]
			prev1 = prev2
			prev2 = nums[i]
			k++
		} else if prev1 != prev2 && nums[i] == prev2 {
			nums[k] = nums[i]
			prev1 = prev2
			prev2 = nums[i]
			k++
		} else if prev1 != prev2 && nums[i] != prev2 {
			nums[k] = nums[i]
			prev1 = prev2
			prev2 = nums[i]
			k++
		} else {
			prev1 = prev2
			prev2 = nums[i]
		}
	}
	// fmt.Printf("Array: %v\n", nums)
	return k
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	ans := removeDuplicatesII(nums)
	fmt.Printf("The answer is: %v\n", ans)
}
