package intervals

import (
	"strconv"
)

// func summaryRanges(nums []int) []string {
// 	if len(nums) == 1 {
// 		return []string{string(nums[0])}
// 	}
// 	ans := []string{}

// 	first, second := 0, 1
// 	r := strconv.Itoa(nums[first])

// 	// fmt.Printf("r: %v\n", strconv.Itoa(nums[0]))
// 	for second < len(nums) {
// 		if nums[second] != nums[second-1]+1 {
// 			if nums[second-1] == nums[first] {
// 				ans = append(ans, r)
// 				r = strconv.Itoa(nums[second])
// 				first = second
// 			} else {
// 				r += "->" + strconv.Itoa(nums[second-1])
// 				ans = append(ans, r)
// 				r = strconv.Itoa(nums[second])
// 				first = second
// 			}

// 			if second == len(nums)-1 {
// 				if second == first {
// 					ans = append(ans, r)
// 				} else {
// 					r += "->" + strconv.Itoa(nums[second])
// 					ans = append(ans, r)
// 				}
// 			}
// 		} else if second == len(nums)-1 {
// 			if second == first {
// 				ans = append(ans, r)
// 			} else {
// 				r += "->" + strconv.Itoa(nums[second])
// 				ans = append(ans, r)
// 			}
// 		}
// 		second++
// 	}

// 	return ans
// }

func summaryRanges(nums []int) []string {
	ans := []string{}
	r := ""

	for i := 0; i < len(nums); i++ {
		r += strconv.Itoa(nums[i])

		j := i + 1
		for j < len(nums) && nums[j]-nums[j-1] == 1 {

			j++
		}

		if j-1 > i {
			r += "->" + strconv.Itoa(nums[j-1])
			ans = append(ans, r)
			r = ""
		} else {
			ans = append(ans, r)
			r = ""
		}

		i = j - 1
	}

	return ans
}

// func main() {
// 	nums := []int{0, 1, 2, 4, 5, 7}

// 	fmt.Printf("Ans: %v\n", summaryRanges(nums))
// }
