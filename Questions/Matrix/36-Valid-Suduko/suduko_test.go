package matrix

import "testing"

func TestIsValidSuduko(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	t.Run("Test valid board", func(t *testing.T) {
		if !isValidSuduko(board) {
			t.Errorf("got: false, want: true")
		}
	})

	t.Run("Test invalid board", func(t *testing.T) {
		board[0][0] = '8'
		if isValidSuduko(board) {
			t.Errorf("got: true, want: want")
		}
	})
}
