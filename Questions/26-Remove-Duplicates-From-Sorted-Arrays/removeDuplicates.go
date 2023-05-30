package dsa

import "fmt"

func removeDuplicates(nums []int) int {
	// initialise map
	prev := nums[0]

	k := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[k] = nums[i]
			k++
		}
		prev = nums[i]
	}
	fmt.Printf("Array without duplicates: %v\n", nums)
	return k
}

// This main function is just used for iteration/testing
func main() {
	nums := []int{1, 2, 2, 3, 3, 4, 4, 5, 5}
	ans := removeDuplicates(nums)

	fmt.Printf("The length of unique elements is: %v\n", ans)
}
