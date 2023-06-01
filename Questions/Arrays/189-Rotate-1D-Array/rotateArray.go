package arrays

import "fmt"

func rotate(nums []int, k int) []int {
	// This is solvable in a 3 step process: (Example Array: [1, 2, 3, 4, 5])
	// 	1. Reverse the entire array
	//		a. Array: [5, 4, 3, 2, 1] this is an O(n)
	//	2. Reverse 0:k
	// 		a. Array: [4, 5, 3, 2, 1] this is an O(k)
	// 	3. Reverse k:n-1
	// 		a. Array: [4, 5, 1, 2, 3] this is an O(n-k)
	// Overall making it an O(n) + O(k) + O(n-k) which for all intents and purposes is O(n)
	// Even in a worse case scenario where k >= n you get O(2n) ~= O(n)

	// To catch cases where k > n we can make our pivot point modular to the length
	pivot := k % len(nums)

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, pivot-1)
	reverse(nums, pivot, len(nums)-1)
	return nums
}

// Go passes by pointers so no need to return nums here
func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}

	ans := rotate(nums, 2)

	fmt.Printf("Answer: %v", ans)
}
