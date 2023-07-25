package graphs

import "testing"

func TestNumIslands(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '1'},
		{'0', '0', '1', '0', '0'},
	}

	t.Run("Test valid grid", func(t *testing.T) {
		want := 3

		got := numIslands(grid)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
