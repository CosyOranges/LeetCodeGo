package arrays

import "fmt"

func majorityElement(nums []int) int {
	// The easy solution is to use a map
	elements := map[int]int{}
	k := len(nums)
	ans := 0

	for i := 0; i < len(nums); i++ {
		elements[nums[i]]++
		if elements[nums[i]] > k/2 {
			ans = nums[i]
		}
	}

	return ans
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	ans := majorityElement(nums)

	fmt.Printf("Answer: %v\n", ans)
}
