package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func (s *ListNode) Build(arr []int) *ListNode {
	if s == nil {
		s = &ListNode{Val: arr[0], Next: nil}
	}

	s.Next = s.buildHelper(arr, 1)

	return s
}

func (s *ListNode) buildHelper(arr []int, ind int) *ListNode {
	if ind > len(arr)-1 {
		return nil
	}
	node := &ListNode{Val: arr[ind], Next: nil}

	node.Next = node.buildHelper(arr, ind+1)

	return node
}

func (s *ListNode) CreateCycle(pos1 int, pos2 int) {
	ind := 1
	first, second := s, s
	for ind < pos1 {
		first = first.Next
		ind++
	}

	for ind <= pos2 {
		second = second.Next
		ind++
	}

	second.Next = first
}

func hasCycle(head *ListNode) bool {
	pos := -1
	if head == nil {
		return false
	}
	posMap := make(map[*ListNode]int)
	pos = 0
	for head.Next != nil {
		if _, ok := posMap[head]; ok {
			return true
		} else {
			posMap[head] = pos
		}
		head = head.Next
		pos++
	}

	return false
}

func hasCycleFloyd(node *ListNode) bool {
	if node == nil {
		return false
	}

	slow, fast := node, node

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

// func main() {
// 	linkedList := &ListNode{}
// 	linkedList = linkedList.Build([]int{3, 2, 0, -4})
// 	linkedList.CreateCycle(2, 4)
// 	fmt.Printf("Has Cycle: %v\n", hasCycleFloyd(linkedList))
// }
