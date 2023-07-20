package segmentTree

import "testing"

func TestSegmentTree(t *testing.T) {
	test := NewTree([]int{9, -8})

	t.Run("Test Update", func(t *testing.T) {
		test.Update(0, 3)
		if test.nodes[1] != 3 {
			t.Errorf("want: 3, got: %v", test.nodes[0])
		}
	})

	t.Run("Test SumRange", func(t *testing.T) {
		x := test.SumRange(1, 1)

		if x != -8 {
			t.Errorf("want: 8, got: %v", x)
		}
	})

	t.Run("Test size", func(t *testing.T) {
		if len(test.nodes) != 3 {
			t.Errorf("Size allocated incorrectly, should be 3, is: %v", len(test.nodes))
		}
	})
}
