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

func (t *BinaryTree[T]) BuildFromArray(arr []T) *BinaryTree[T] {
	n := len(arr)
	if n == 0 {
		return nil
	}

	ind := 0
	t.root = t.root.newNode(arr[ind])
	if ind < n {
		t.root.left = t.root.buildFromArray(arr, 2*ind+1)
		t.root.right = t.root.buildFromArray(arr, 2*ind+2)
	}
	return t
}

func (b *binarynode[T]) buildFromArray(arr []T, ind int) *binarynode[T] {

	node := b.newNode(arr[ind])
	if ind < len(arr) {
		node.left = b.buildFromArray(arr, 2*ind+1)
		node.right = b.buildFromArray(arr, 2*ind+2)
	} else {
		return nil
	}

	return node
}

func (t *BinaryTree[T]) BuildFromInAndPostOrderLists(InOrder []T, PostOrder []T) *BinaryTree[T] {
	n := len(PostOrder)
	// If the current node has no children return
	if n == 0 {
		return nil
	}

	t.root = t.root.newNode(PostOrder[n-1])

	// If PostOrder list is only of length 1 then we return the root node
	if n <= 1 {
		return t
	}

	// Find the index in `InOrder` about which to split
	split_indx := findSplitIndex(InOrder, PostOrder[n-1])

	// Create the Left and Right InOrder slices
	left_inorder_slice := InOrder[:split_indx]
	right_inorder_slice := InOrder[split_indx+1:]

	// Create the Left and Right PostOrder slices:
	// Left will be from index 0 -> len(left_inorder_slice)
	// Right will be from index len(left_inorder_slice) -> n-1 (not inclusive becuase PostOrder[n-1] is the root)
	left_postorder_slice := PostOrder[:len(left_inorder_slice)]
	right_postorder_slice := PostOrder[len(left_inorder_slice) : n-1]

	// Recursively build the rest of the tree, this helper function follows the same logic, but returns a BinaryNode instead of the tree
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

	t.root.walkPreOrder(t.root.left, arr, i)
	t.root.walkPreOrder(t.root.right, arr, i)

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

// Priivate functions to the package, these are helper functions to the above
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

func (b *binarynode[T]) walkPreOrder(node *binarynode[T], arr []T, i *int) {
	if node == nil {
		return
	}
	arr[*i] = node.val
	*i++
	b.walkPreOrder(node.left, arr, i)
	b.walkPreOrder(node.right, arr, i)
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
