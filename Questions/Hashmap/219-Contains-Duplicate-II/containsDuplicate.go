package hashmap

import (
	"fmt"
	"math"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	iMap := make(map[int]int)

	for i, x := range nums {
		if _, ok := iMap[x]; !ok {
			iMap[x] = i
		} else {
			if int(math.Abs(float64(iMap[x]-i))) <= k {
				return true
			} else {
				iMap[x] = i
			}
		}
	}

	return false
}

func main() {
	fmt.Printf("Ans: %v\n", containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 3))
}
