package arrays

import "fmt"

func ZigZag(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	ans := ""
	grid := make([][]byte, len(s))

	for i := 0; i < len(grid); i++ {
		grid[i] = make([]byte, numRows)
	}

	wordInd := 0
	zigzag := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < numRows; j++ {
			if i%(numRows-1) == 0 && wordInd < len(s) {
				grid[i][j] = s[wordInd]
				wordInd++
				zigzag = numRows - 2
			} else if j == zigzag && wordInd < len(s) {
				grid[i][j] = s[wordInd]
				wordInd++
				zigzag -= 1
			} else {
				grid[i][j] = '#'
			}
			fmt.Printf("%c ", grid[i][j])
		}
		fmt.Print("\n")
	}

	for j := 0; j < numRows; j++ {
		for i := 0; i < len(grid); i++ {
			if grid[i][j] != '#' {
				ans += string(grid[i][j])
				if len(ans) == len(s) {
					return ans
				}
			}
		}
	}

	return ans
}

func main() {
	fmt.Printf("ANS: %v\n", ZigZag("TESTINGTHISTEST", 4))
}
