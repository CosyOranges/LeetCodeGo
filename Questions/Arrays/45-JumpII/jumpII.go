package arrays

import "fmt"

func jump(nums []int) int {
	curJump, jumps, farthest := 0, 0, 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+i > farthest {
			farthest = nums[i] + i
		}

		if curJump == i {
			curJump, jumps = farthest, jumps+1

			if curJump >= len(nums)-1 {
				return jumps
			}
		}
	}

	return 0
}

func main() {
	nums := []int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}

	ans := jump(nums)
	fmt.Printf("Ans: %v\n", ans)
}
