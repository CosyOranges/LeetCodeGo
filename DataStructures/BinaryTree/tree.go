package binarytree

import (
	"reflect"
)

type binarynode[T any] struct {
	val   T
	left  *binarynode[T]
	right *binarynode[T]
}

type BinaryTree[T any] struct {
	root *binarynode[T]
}

func (t *BinaryTree[T]) BuildFromInAndPostOrderLists(InOrder []T, PostOrder []T) *BinaryTree[T] {
	n := len(PostOrder)
	// If the current node has no children return
	if n == 0 {
		return nil
	}

	t.root = t.root.newNode(PostOrder[n-1])

	if n <= 1 {
		return t
	}

	split_indx := findSplitIndex(InOrder, PostOrder[n-1])

	left_inorder_slice := InOrder[:split_indx]
	right_inorder_slice := InOrder[split_indx+1:]

	left_postorder_slice := PostOrder[:len(left_inorder_slice)]
	right_postorder_slice := PostOrder[len(left_inorder_slice) : len(PostOrder)-1]

	t.root.left = t.root.buildFromInAndPostOrderLists(left_inorder_slice, left_postorder_slice)
	t.root.right = t.root.buildFromInAndPostOrderLists(right_inorder_slice, right_postorder_slice)

	return t
}

func (t *BinaryTree[T]) WalkPreOrder() []T {
	if t == nil {
		return nil
	}
	size := t.Size()
	arr := make([]T, size)

	// Here we are creating our index as a pointer so we don't need to return anything
	var i *int
	i = new(int)
	*i = 0
	arr[*i] = t.root.val
	*i++

	t.root.walkPreOrderHelper(t.root.left, arr, i)
	t.root.walkPreOrderHelper(t.root.right, arr, i)

	return arr
}

func (t *BinaryTree[T]) Size() int {
	if t.root == nil {
		return 0
	}

	size := 1
	size += t.root.size(t.root.left)
	size += t.root.size(t.root.right)

	return size
}

func (b *binarynode[T]) buildFromInAndPostOrderLists(InOrder []T, PostOrder []T) *binarynode[T] {
	n := len(PostOrder)
	// If the current node has no children return
	if n == 0 {
		return nil
	}

	node := b.newNode(PostOrder[n-1])

	if n <= 1 {
		return node
	}

	split_indx := findSplitIndex(InOrder, PostOrder[n-1])

	left_inorder_slice := InOrder[:split_indx]
	right_inorder_slice := InOrder[split_indx+1:]

	left_postorder_slice := PostOrder[:len(left_inorder_slice)]
	right_postorder_slice := PostOrder[len(left_inorder_slice) : len(PostOrder)-1]

	node.left = b.buildFromInAndPostOrderLists(left_inorder_slice, left_postorder_slice)
	node.right = b.buildFromInAndPostOrderLists(right_inorder_slice, right_postorder_slice)

	return node
}

func findSplitIndex[T any](InOrder []T, val T) int {
	for i, x := range InOrder {
		if reflect.DeepEqual(x, val) {
			return i
		}
	}

	return -1
}

func (b *binarynode[T]) walkPreOrderHelper(node *binarynode[T], arr []T, i *int) {
	if node == nil {
		return
	}
	arr[*i] = node.val
	*i++
	b.walkPreOrderHelper(node.left, arr, i)
	b.walkPreOrderHelper(node.right, arr, i)
}

func (b *binarynode[T]) size(node *binarynode[T]) int {
	if node == nil {
		return 0
	}

	size := 1

	size += b.size(node.left)
	size += b.size(node.right)

	return size
}

func (b *binarynode[T]) newNode(val T) *binarynode[T] {
	node := &binarynode[T]{
		val:   val,
		left:  nil,
		right: nil,
	}
	return node
}

// func main() {
// 	tree := &BinaryTree[int]{}
// 	// tree2 := &BinaryTree[int]{}
// 	InOrder := []int{4, 8, 2, 5, 1, 6, 3, 7}
// 	PostOrder := []int{8, 4, 5, 2, 6, 7, 3, 1}
// 	// PreOrder := []int{1, 2, 4, 8, 5, 3, 6, 7}
// 	tree.BuildFromInAndPostOrderLists(InOrder, PostOrder)
// 	PreOrder := tree.WalkPreOrder()
// 	fmt.Printf("size: %v\n", tree.Size())
// 	fmt.Printf("PreOrder: %v\n", PreOrder)
// }
