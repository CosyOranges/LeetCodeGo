package arrays

import (
	"fmt"
	"sort"
)

func threesum(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)

	// Now we can loop through and ignore duplicates
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			target := nums[i] + nums[left] + nums[right]

			if target == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				// Now increment left and right until they no longer are the same as indices next to them
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if target < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return ans
}

func main() {
	test := []int{-1, 0, 1, 2, -1, -4}

	sol := threesum(test)

	for _, x := range sol {
		fmt.Printf("Solution Triplet: %v\n", x)
	}
}
