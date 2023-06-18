package linkedlist

import (
	linkedlist "dsa/DataStructures/LinkedList"
)

func addTwoNumbers(l1 *linkedlist.ListNode, l2 *linkedlist.ListNode) *linkedlist.ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	return addTwoNumbersHelp(l1, l2, 0)
}

func addTwoNumbersHelp(l1 *linkedlist.ListNode, l2 *linkedlist.ListNode, carry int) *linkedlist.ListNode {
	sum := carry
	res := 0
	if l1 == nil && l2 != nil {
		sum += l2.Val
		if sum >= 10 {
			res = sum - 10
		} else {
			res = sum
		}
		l2.Val = res
		l2.Next = addTwoNumbersHelp(nil, l2.Next, (sum-res)/10)
		return l2
	} else if l2 == nil && l1 != nil {
		sum += l1.Val
		if sum >= 10 {
			res = sum - 10
		} else {
			res = sum
		}
		l1.Val = res
		l1.Next = addTwoNumbersHelp(l1.Next, nil, (sum-res)/10)
		return l1
	} else if l1 == nil && l2 == nil && sum != 0 {
		node := &linkedlist.ListNode{Val: sum, Next: nil}
		return node
	} else if l1 == nil && l2 == nil && sum == 0 {
		return nil
	} else {
		sum += l1.Val + l2.Val
	}

	if sum >= 10 {
		res = sum - 10
	} else {
		res = sum
	}
	carry = (sum - res) / 10
	node := &linkedlist.ListNode{Val: res, Next: nil}
	node.Next = addTwoNumbersHelp(l1.Next, l2.Next, carry)

	return node
}

func main() {
	list1 := &linkedlist.ListNode{}
	list1 = list1.Build([]int{9, 9, 9, 9, 9, 9, 9})
	list2 := &linkedlist.ListNode{}
	list2 = list2.Build([]int{9, 9, 9, 9})

	added := addTwoNumbers(list1, list2)
	added.Show()
	// fmt.Printf("Added: %v\n", list1)
}
