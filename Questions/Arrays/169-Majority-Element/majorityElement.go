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

func majorityElementConstantSpace(nums []int) int {
	// To achieve this without using a map you can use the Boyer-Moore Majority Vote algorithm
	// This works by:
	// Assigning a candidate to the first position of the array
	candidate := nums[0]

	// Assigning a count to keep track of the instances of the candidate
	count := 1

	for _, num := range nums[1:] {
		if num == candidate {
			count++
		} else {
			count--
		}
		// If the count every reaches 0 then the new candidate is our current index
		// Which every number is repeated most frequently will be the final candidate
		// Provided that frequency is greater than len(nums) / 2
		if count == 0 {
			candidate = num
			count = 1
		}
	}

	return candidate
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	ans := majorityElement(nums)

	fmt.Printf("Answer: %v\n", ans)
}
