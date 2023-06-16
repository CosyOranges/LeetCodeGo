package arrays

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	area := 0

	for left < right {
		temp := (right - left) * min(height[left], height[right])
		if temp > area {
			area = temp
		} else if height[left] < height[right] && height[left+1] > height[left] {
			left++
		} else if height[right] < height[left] && height[right-1] > height[right] {
			right--
		} else {
			if height[left] < height[right] {
				left++
			} else {
				right--
			}
		}
	}

	return area
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	height := []int{1, 1}

	fmt.Printf("Ans: %v\n", maxArea(height))
}
