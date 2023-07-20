package segmentTree

import (
	"fmt"
	"math"
)

type SegmentTree struct {
	nodes []int // Nodes to store info in
	size  int   // size of the original array
}

func NewTree(nums []int) *SegmentTree {
	// treeSize := calcTreeSize(len(nums))
	nodes := make([]int, calcTreeSize(len(nums)))

	t := &SegmentTree{nodes: nodes, size: len(nums)}
	t.build(nums, 0, 0, len(nums)-1)

	return t
}

func (t *SegmentTree) build(nums []int, node int, left int, right int) {
	if left > right {
		return
	} else if left == right {
		// fmt.Printf("HERE, left: %v | right: %v\n", left, right)
		t.nodes[node] = nums[left]
	} else {
		mid := left + (right-left)/2

		t.build(nums, 2*node+1, left, mid)
		t.build(nums, 2*node+2, mid+1, right)

		t.nodes[node] += t.nodes[2*node+1] + t.nodes[2*node+2]
	}
}

func (t *SegmentTree) Update(x, y int) {
	t.update(0, 0, t.size-1, x, y)
}

func (t *SegmentTree) update(node int, left int, right int, updateIndex int, val int) {
	if left > right || left > updateIndex || right < updateIndex {
		return
	} else if left == right {
		t.nodes[node] = val
	} else {
		mid := (left + right) / 2

		t.update(2*node+1, left, mid, updateIndex, val)
		t.update(2*node+2, mid+1, right, updateIndex, val)
		t.nodes[node] = t.nodes[2*node+1] + t.nodes[2*node+2]
	}
}

func (t *SegmentTree) SumRange(queryL int, queryR int) int {
	return t.sumRange(0, 0, t.size-1, queryL, queryR)
}

func (t *SegmentTree) sumRange(node int, left int, right int, queryL int, queryR int) int {
	if queryL <= left && queryR >= right {
		fmt.Printf("Left: %v query: %v | Right: %v query: %v\n", left, queryL, right, queryR)
		return t.nodes[node]
	}

	if queryL > right || queryR < left {
		return 0
	}

	mid := left + (right-left)/2

	return t.sumRange(2*node+1, left, mid, queryL, queryR) + t.sumRange(2*node+2, mid+1, right, queryL, queryR)
}

func calcTreeSize(originalSize int) int {
	x := math.Ceil(math.Log2(float64(originalSize)))

	return 2*int(math.Pow(2, x)) - 1
}

// func main() {
// 	// fmt.Printf("Size: %v\n", calcTreeSize(6))
// 	tree := NewTree([]int{9, -8})
// 	fmt.Printf("Test Tree: %v\n", tree.nodes)
// 	tree.Update(0, 3)
// 	fmt.Printf("Test Tree: %v\n", tree.nodes)
// 	fmt.Printf("Sum: %v\n", tree.SumRange(1, 1))
// }
