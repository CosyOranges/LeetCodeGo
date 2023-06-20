package dynamic

import "fmt"

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] = min(grid[i][j]+dp[i-1][j], grid[i][j]+dp[i][j-1])
			}
		}
	}
	PrintGrid(dp)
	return dp[len(dp)-1][len(dp[len(dp)-1])-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func PrintGrid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if j == 0 {
				fmt.Printf("[ %v", grid[i][j])
			} else if j == len(grid[i])-1 {
				fmt.Printf("%v ]", grid[i][j])
			} else {
				fmt.Printf(" %v ", grid[i][j])
			}
		}
		fmt.Println()
	}
}

func main() {
	// grid := [][]int{
	// 	{1, 3, 1},
	// 	{1, 5, 1},
	// 	{4, 2, 1},
	// }
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	PrintGrid(grid)
	ans := minPathSum(grid)
	fmt.Printf("Ans: %v\n", ans)
}
