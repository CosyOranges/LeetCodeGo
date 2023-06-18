package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (s *ListNode) Build(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	s = &ListNode{Val: arr[0], Next: nil}

	s.Next = s.buildHelper(arr, 1)

	return s
}

func (s *ListNode) Show() {
	if s == nil {
		fmt.Println("[Empty]")
		return
	}

	curr := s
	fmt.Print("[")
	for curr.Next != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
	fmt.Printf("%d]\n", curr.Val)
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
