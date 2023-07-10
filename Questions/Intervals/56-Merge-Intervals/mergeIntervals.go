package intervals

import (
	"fmt"
	"sort"
)

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// Sort the rows of the intervals based on the first column
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Create answer array
	ans := [][]int{}

	// Create a left and a row pointer
	anchor, tMax, frontRow := intervals[0][0], intervals[0][1], 1
	// currMin, currMax := intervals[1][0], intervals[1][1]
	for frontRow < len(intervals) {
		if intervals[frontRow][0] <= tMax {
			// We can expeand our interval
			if intervals[frontRow][1] > tMax {
				tMax = intervals[frontRow][1]
			}
		} else if intervals[frontRow][0] > tMax {
			// Otherwise the current index is not valid in the built interval
			// So create the interval with our anchor and tMax
			ans = append(ans, []int{anchor, tMax})

			// Reset all values to the current
			anchor, tMax = intervals[frontRow][0], intervals[frontRow][1]
		}
		frontRow++
	}

	// One final check
	if intervals[frontRow-1][1] > tMax {
		tMax = intervals[frontRow-1][1]
	}
	ans = append(ans, []int{anchor, tMax})

	return ans
}

func main() {
	// [][]int{{1, 4}, {2, 3}, {8, 10}, {9, 20}}
	// {2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}
	fmt.Printf("Ans: %v\n", mergeIntervals([][]int{{2, 3}, {4, 6}, {5, 7}, {3, 4}}))
}
