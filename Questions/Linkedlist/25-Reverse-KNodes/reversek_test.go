package linkedlist

import (
	linkedlist "dsa/DataStructures/LinkedList"
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	list1 := &linkedlist.ListNode{}
	list1 = list1.Build([]int{1, 2, 3, 4, 5})

	t.Run("Test with k = 1", func(t *testing.T) {
		want := &linkedlist.ListNode{}
		want = want.Build([]int{1, 2, 3, 4, 5})

		got := reverseKGroup(list1, 1)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("List was not left unchanged")
		}
	})

	t.Run("Test with k = 2", func(t *testing.T) {
		want := &linkedlist.ListNode{}
		want = want.Build([]int{2, 1, 4, 3, 5})

		got := reverseKGroup(list1, 2)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("List was not [2, 1, 4, 3, 5]")
		}
	})
}
