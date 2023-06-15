package arrays

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	ans := ""
	if num == 0 {
		return ans
	}

	numerals := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for _, x := range nums {
		if num >= x {
			// Dividing an integer will give floor(x/y) e.g. 3999/1000 = 3
			residual := num / x
			fmt.Printf("Res: %v | num: %v | x: %v\n", residual, num, x)
			// We then want to repeat numerals[x] residual times
			// e.g. for 3999/1000 = 3, we want to repeat 1000(M) 3 times
			// for situations like 900, 90, or 9 the residual will only ever be 1
			ans += strings.Repeat(numerals[x], residual)

			// Finally we take the mod of num % x ==> 3999 % 1000 = 999
			num = num % x
		}
	}
	return ans
}

func intToSlice(num int, arr []int) []int {
	if num != 0 {
		i := num % 10
		arr = append([]int{i}, arr...)
		return intToSlice(num/10, arr)
	}

	return arr
}

// func main() {
// 	// 3999 -> MMMCMXCIX
// 	fmt.Printf("Ans: %s\n", intToRoman(3999))
// 	// fmt.Printf("Mod 4 % 1000: %v\n")
// }
