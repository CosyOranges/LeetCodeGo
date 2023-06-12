package arrays

import "fmt"

func trap(height []int) int {
	// A first pass solution attempt:
	// valley, valleyHeight := false, 0
	// valleys, depths := make([]bool, len(height)), make([]int, len(height))
	// for i := 0; i < len(height); i++ {
	// 	if i > 0 && !valley && height[i] < height[i-1] {
	// 		valley = true
	// 		valleyHeight = height[i-1]
	// 	} else if valley && height[i] > valleyHeight {
	// 		valley = false
	// 		valleyHeight = height[i]
	// 	}

	// 	if valley {
	// 		depths[i] = valleyHeight - height[i]
	// 	}
	// 	valleys[i] = valley
	// }

	// valley, valleyHeight = false, 0
	// for i := len(height) - 1; i >= 0; i-- {
	// 	if i < len(height)-1 && !valley && height[i] < height[i+1] {
	// 		valley = true
	// 		valleyHeight = height[i+1]
	// 	} else if valley && height[i] >= valleyHeight {
	// 		valley = false
	// 	}
	// 	if valleyHeight-height[i] < depths[i] {
	// 		depths[i] = valleyHeight - height[i]
	// 	}
	// 	if valleys[i] == valley {
	// 		valleys[i] = valley
	// 	} else {
	// 		valleys[i] = false
	// 	}
	// }

	// valleyHeight = 0
	// volume := 0
	// for i := 0; i < len(height); i++ {
	// 	if depths[i] > 0 {
	// 		volume += depths[i]
	// 	}
	// }

	// A more optimal solution can be done using Two pointers from either
	// end of the array
	left, right := 0, len(height)-1
	maxLeftHeight, maxRightHeight, volume := 0, 0, 0

	for left < right {
		if height[left] < height[right] {
			if height[left] > maxLeftHeight {
				maxLeftHeight = height[left]
			} else {
				volume += maxLeftHeight - height[left]
			}
			left++
		} else {
			if height[right] > maxRightHeight {
				maxRightHeight = height[right]
			} else {
				volume += maxRightHeight - height[right]
			}
			right--
		}
	}

	return volume
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	fmt.Printf("Volume: %v\n", trap(height))
}
