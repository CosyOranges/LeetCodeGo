package matrix

func isValidSuduko(board [][]byte) bool {
	// We need 3 maps: (1) Track the 9 by 1 rows, (2) Track the 1 by 9 columns, (3) Track the 3 by 3 boxes
	// In the rows map [] is the row number, [][] is the digit in that row, the first time we encounter it we set it to true
	// if we ever encounter it again in the same row then we return false as the Suduko board is invalid
	// The Boxes map is slightly different [] represents the box, to find which box we are in we need to: [i/3][j/3][]
	// because 1/3 rounds to 1 in integer math
	columns, rows, boxes := [9][9]bool{}, [9][9]bool{}, [3][3][9]bool{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			cell := board[i][j]

			// If the cell hasn't been given a digit then continue
			if cell == '.' {
				continue
			}
			cell -= '1'
			// Check if the value already exists in the row, column, or box
			if columns[j][cell] || rows[i][cell] || boxes[i/3][j/3][cell] {
				return false
			} else {
				columns[j][cell] = true
				rows[i][cell] = true
				boxes[i/3][j/3][cell] = true
			}
		}
	}

	// If we made it here without returning false, the Suduko board must be valid
	return true
}

func main() {

}
