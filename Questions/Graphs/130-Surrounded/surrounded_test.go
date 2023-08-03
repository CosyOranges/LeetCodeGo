package graphs

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Run("Test empty board", func(t *testing.T) {
		var board [][]byte = [][]byte{}
		solve(board)

		if !reflect.DeepEqual(board, [][]byte{}) {
			t.Errorf("got: %v, want: %v", board, [][]byte{})
		}
	})

	t.Run("Test actual board", func(t *testing.T) {
		board := [][]byte{
			{'X', 'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X', 'X'},
			{'X', 'O', 'O', 'X', 'X'},
			{'X', 'X', 'O', 'X', 'X'},
			{'X', 'X', 'X', 'O', 'O'},
			{'X', 'X', 'O', 'X', 'X'},
		}

		want := [][]byte{
			{'X', 'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'O', 'O'},
			{'X', 'X', 'O', 'X', 'X'},
		}

		solve(board)

		if !reflect.DeepEqual(want, board) {
			t.Errorf("got: %v, want: %v", board, want)
		}
	})
}
