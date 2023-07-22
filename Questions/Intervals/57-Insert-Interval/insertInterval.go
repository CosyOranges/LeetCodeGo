package intervals

import (
	"fmt"
	"sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	ans := [][]int{}

	// append and sort
	intervals = append(intervals, newInterval)

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	anchor, front, next := intervals[0][0], intervals[0][1], 1

	for next < len(intervals) {
		if intervals[next][0] <= front && intervals[next][1] > front {
			front = intervals[next][1]
		} else if intervals[next][0] > front {
			ans = append(ans, []int{anchor, front})
			anchor, front = intervals[next][0], intervals[next][1]
		}
		next++
	}

	ans = append(ans, []int{anchor, front})

	return ans
}

func main() {
	fmt.Printf("Ans: %v\n", insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}
