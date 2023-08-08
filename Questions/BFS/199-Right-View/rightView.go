package bfs

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ans []int = []int{}
	var queue []*TreeNode = []*TreeNode{root}

	for len(queue) > 0 {
		var rowLength int = len(queue)
		var i int = 0

		for i < rowLength {
			// Add left and right children if present
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			i++
		}

		// Add the most right node to our answer
		ans = append(ans, queue[i-1].Val)
		queue = queue[rowLength:]
	}

	return ans
}

// func main() {
// 	tree := &TreeNode{Val: 1}
// 	tree.Left = &TreeNode{Val: 2}
// 	tree.Left.Right = &TreeNode{Val: 5}
// 	tree.Right = &TreeNode{Val: 3}
// 	tree.Right.Right = &TreeNode{Val: 4}
// 	tree.Right.Right.Left = &TreeNode{Val: 20}

// 	fmt.Printf("Solution: %v\n", rightSideView(tree))
// }
