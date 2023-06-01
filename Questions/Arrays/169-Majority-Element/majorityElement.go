package arrays

import "fmt"

func majorityElement(nums []int) int {
	// The easy solution is to use a map
	elements := map[int]int{}
	k := len(nums)

	for i := 0; i < len(nums); i++ {
		if _, ok := elements[nums[i]]; ok {
			elements[nums[i]]++
		} else {
			elements[nums[i]] = 1
		}
	}

	ans := 0
	for key, val := range elements {
		if val > k/2 {
			ans = key
		}
	}

	return ans
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	ans := majorityElement(nums)

	fmt.Printf("Answer: %v\n", ans)
}
