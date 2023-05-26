package dsa

import "fmt"

func mergeArrays(nums1 []int, m int, nums2 []int, n int) []int {
	if n != 0 && m != 0 {
		// First shift all of nums1 to the end of the array in sorted order
		j := len(nums1) - 1
		for i := m - 1; i >= 0; i-- {
			nums1[j] = nums1[i]
			j--
		}

		fmt.Printf("The shifted nums1 array is: %d.\n", nums1)
		// Now we can compare with nums2 and insert the smaller of the numbers at the start of the list
		i := n // current index in nums1
		j = 0  // current index in nums2
		for k := 0; k < m+n; k++ {
			if i < m+n && j < n {
				if nums1[i] < nums2[j] {
					nums1[k] = nums1[i]
					i++
				} else {
					nums1[k] = nums2[j]
					j++
				}
			} else if j < n {
				nums1[k] = nums2[j]
				j++
			}
		}

	} else if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
	}

	return nums1
}

func main() {
	nums1 := []int{4, 0, 0, 0, 0, 0}
	nums2 := []int{1, 2, 3, 5, 6}
	nums1 = mergeArrays(nums1, 1, nums2, 5)

	fmt.Printf("The final nums1 array: %d.\n", nums1)
}
