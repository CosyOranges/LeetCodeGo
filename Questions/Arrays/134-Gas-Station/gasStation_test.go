package arrays

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	gas, cost := []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}

	t.Run("Standard test", func(t *testing.T) {
		want := 3

		got := canCompleteCircuit(gas, cost)

		if want != got {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
