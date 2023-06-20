package dynamic

import "testing"

func TestMinPathSum(t *testing.T) {
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	grid2 := [][]int{
		{1},
	}

	t.Run("Test standard array", func(t *testing.T) {
		want := 12

		got := minPathSum(grid)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test smallest array", func(t *testing.T) {
		want := 1

		got := minPathSum(grid2)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
