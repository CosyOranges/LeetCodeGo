package arrays

import "fmt"

func canJump(nums []int) bool {
	remaining := len(nums)

	jumps_left := 0

	for i := 0; i < len(nums); i++ {
		jumps_left--
		remaining--
		if jumps_left <= 0 && nums[i] == 0 && remaining > 0 {
			return false
		} else if nums[i] > jumps_left {
			jumps_left = nums[i]
		}

		fmt.Printf("Jumps Left: %v | Remaining: %v\n", jumps_left, remaining)
	}

	fmt.Println("Here")
	return jumps_left >= remaining
}

func main() {
	// nums := []int{2, 3, 1, 1, 4}
	nums2 := []int{0}

	ans := canJump(nums2)
	fmt.Printf("Answer: %v\n", ans)
}
