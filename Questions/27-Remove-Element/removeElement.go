package dsa

import "fmt"

func removeElement(nums []int, val int) int {
	// initialise a count and if the array is empty return 0
	i := 0
	for k := 0; k < len(nums); {
		if nums[k] != val {
			nums[i] = nums[k]
			i++
		}
		k++
	}
	fmt.Printf("The final array looks like: %v", nums)
	return len(nums[0:i])
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	ans := removeElement(nums, val)
	fmt.Printf("The number of %v in nums is: %v", val, ans)
}
