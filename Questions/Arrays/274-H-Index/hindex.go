package arrays

import (
	"fmt"
	"sort"
)

func hIndex(nums []int) int {
	// The maximum h index a researcher can have is len(nums)

	// Every paper that is not cited at least hmax times, reduced the hmax by 1
	// If hmax comes down enough then it is possible that previous papers are now valid for the hmax
	// e.g. [3, 0, 6, 1, 5]
	// hmax = 6
	// [i = 0] 1st paper is cited 3 times, potential hIndex i 1
	// [i = 1] 2nd paper is cited 0 times, potential hIndex is 2, but current hIndex is 1
	// [i = 2] 3rd paper is cited 6 times, potential hIndex is 3, but current hIndex is 2
	// [i = 3] 4th paper is cited 1 times, potential hIndex is 4, but current hIndex is 2
	// [i = 4] 5th paper is cited 5 times, potential hIndex is 5, but current hIndex is 3
	/*
		The simplest way to solve this problem is to first sort the array
		sort(nums) => [0, 1, 3, 5, 6]

		Then you can iterate backwards bearing in mind that hmax = len(nums) [5 in this case]
		we can track a count (this represents the potential number of citations) and an ans variable
			if nums[i] >= count then ans = count
			count++
	*/

	sort.Ints(nums)
	ans, count := 0, 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] >= count {
			ans = count
		}
		count++
	}

	return ans
}

func main() {
	nums := []int{3, 0, 6, 1, 5}

	fmt.Printf("Ans: %v\n", hIndex(nums))
}
