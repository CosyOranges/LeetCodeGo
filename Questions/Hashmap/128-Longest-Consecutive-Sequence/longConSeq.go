package hashmap

import (
	"fmt"
	"sort"
)

func longestConsecutiveSequence(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	ans := 1
	sort.Ints(nums)
	currSequence := 1
	for i := 1; i < len(nums); i++ {

		if nums[i] == nums[i-1]+1 {
			currSequence++
			if currSequence > ans {
				ans = currSequence
			}
		} else if nums[i] == nums[i-1] {
			continue
		} else {
			currSequence = 1
		}
	}

	return ans
}

func longestConsecutiveSequenceMap(nums []int) int {
	numMap := make(map[int]struct{})

	for _, x := range nums {
		numMap[x] = struct{}{}
	}

	maxLength := 0
	for num := range numMap {
		// Check if current num is the start of a set
		if _, ok := numMap[num-1]; !ok {
			currentSeqLength := 1
			for {
				if _, ok := numMap[num+currentSeqLength]; ok {
					currentSeqLength++
					continue
				}

				if maxLength < currentSeqLength {
					maxLength = currentSeqLength
				}
				break
			}
		}
	}

	return maxLength
}

func main() {
	nums := []int{1, 2, 0, 1}
	fmt.Printf("Ans: %v\n", longestConsecutiveSequence(nums))
}
