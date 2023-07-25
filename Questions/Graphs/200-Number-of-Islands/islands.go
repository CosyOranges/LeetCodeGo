package graphs

import "fmt"

func numIslands(grid [][]byte) int {
	count := 0
	/*
		To solve we will need to loop through every item in the grid and perform a BFS on it if it has not already been visited
		We will mark visited islands with a 2, so we know that they have already been visited
		that way we can increment the count of number of islands everytime we encounter a new 1
	*/
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				bfs(grid, i, j)
			}
		}
	}

	return count
}

func bfs(grid [][]byte, i, j int) {
	// coordinate queue
	queue := [][2]int{}
	queue = append(queue, [2]int{i, j})

	// mark the current coordinates as visited
	grid[i][j] = '2'

	// Create a fixed matrix of neighbour offset
	neighbours := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for len(queue) > 0 {
		currentCoordinates := queue[0]

		// Pop the first set of coordinates from the queue
		queue = queue[1:]

		// Now we need to check the four vertices around the current coordinates
		for _, offset := range neighbours {
			x := currentCoordinates[0] + offset[0]
			y := currentCoordinates[1] + offset[1]

			// check if the neighbours are in the grid
			if x < len(grid) && x >= 0 && y < len(grid[0]) && y >= 0 {
				// if they are and are '1' add them to the queue and set them to 2
				if grid[x][y] == '1' {
					queue = append(queue, [2]int{x, y})
					grid[x][y] = '2'
				}
			}
		}
	}
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '1'},
		{'0', '0', '1', '0', '0'},
	}

	fmt.Printf("Number of Islands: %v\n", numIslands(grid))
}
