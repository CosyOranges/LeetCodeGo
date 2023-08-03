package graphs

import "fmt"

func solve(board [][]byte) {
	// Initial visited board to false
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if !visited[i][j] {
				visited[i][j] = true
				if board[i][j] == 'O' {
					bfs(&board, &visited, i, j)
				}
			}
		}
	}
}

func bfs(board *[][]byte, visited *[][]bool, i, j int) {
	// Coord queue
	queue := [][2]int{}
	queue = append(queue, [2]int{i, j})

	// mark the current as true
	(*visited)[i][j] = true

	// Matrix of neighbours to visit
	var neighbours [4][2]int = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var toFlip bool = true
	var coordsToFlip [][2]int = [][2]int{}

	// Check all neighbours
	for len(queue) > 0 {
		// get current coordinates
		currentCoord := queue[0]
		coordsToFlip = append(coordsToFlip, currentCoord)
		queue = queue[1:]

		for _, offset := range neighbours {
			x := currentCoord[0] + offset[0]
			y := currentCoord[1] + offset[1]

			// Check if the neighbour is a 'O'
			if x < len(*board) && x >= 0 && y < len((*board)[0]) && y >= 0 {
				if (*board)[x][y] == 'O' && !(*visited)[x][y] {
					queue = append(queue, [2]int{x, y})
					(*visited)[x][y] = true
				}
			} else {
				// Otherwise our current 'O' is on the boarder, and therefore we should not flip
				toFlip = false
			}
		}
	}

	// If we should flip then do so
	if toFlip {
		for _, coords := range coordsToFlip {
			(*board)[coords[0]][coords[1]] = 'X'
		}
	}
}

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X', 'X'},
		{'X', 'O', 'O', 'X', 'X'},
		{'X', 'X', 'O', 'X', 'X'},
		{'X', 'X', 'X', 'O', 'O'},
		{'X', 'X', 'O', 'X', 'X'},
	}

	solve(board)
	for _, x := range board {
		fmt.Printf("%v\n", x)
	}
}
