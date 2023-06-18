package linkedlist

import (
	linkedlist "dsa/DataStructures/LinkedList"
)

func mergeTwoLists(list1 *linkedlist.ListNode, list2 *linkedlist.ListNode) *linkedlist.ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1 == nil && list2 == nil {
		return nil
	}

	node := &linkedlist.ListNode{}
	if list1.Val < list2.Val {
		node = list1
		node.Next = mergeTwoLists(list1.Next, list2)
	} else {
		node = list2
		node.Next = mergeTwoLists(list1, list2.Next)
	}

	return node
}

func main() {
	list1 := &linkedlist.ListNode{}
	list1 = list1.Build([]int{1, 2, 4})
	list2 := &linkedlist.ListNode{}
	list2 = list2.Build([]int{1, 3, 4, 5, 6})

	list3 := mergeTwoLists(list1, list2)
	list3.Show()
}
