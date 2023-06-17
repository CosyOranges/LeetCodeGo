package hashmap

import "fmt"

func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int)

	ans := make([]int, 2)
	ans[0], ans[1] = 0, 0
	for i, x := range nums {
		if _, ok := sumMap[target-x]; !ok {
			sumMap[x] = i
		} else if x == target-x && i != sumMap[i] {
			ans[0], ans[1] = i, sumMap[target-x]
		} else {
			ans[0], ans[1] = i, sumMap[target-x]
		}
	}
	return ans
}

func main() {
	nums := []int{2, 2, 7, 7, 11, 15}
	fmt.Printf("Ans: %v\n", twoSum(nums, 9))
}
