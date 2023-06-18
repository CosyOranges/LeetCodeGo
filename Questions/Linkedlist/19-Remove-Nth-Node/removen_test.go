package linkedlist

import (
	linkedlist "dsa/DataStructures/LinkedList"
	"reflect"
	"testing"
)

func TestRemoveNthNode(t *testing.T) {
	list1 := &linkedlist.ListNode{}
	list1 = list1.Build([]int{1})
	list2 := &linkedlist.ListNode{}
	list2 = list2.Build([]int{1, 2, 3, 4, 5})

	t.Run("Test length of 1", func(t *testing.T) {
		got := removeNthNode(list1, 1)

		if got != nil {
			t.Errorf("Did not get empty list")
			got.Show()
		}
	})

	t.Run("Test full list", func(t *testing.T) {
		want := &linkedlist.ListNode{}
		want = want.Build([]int{1, 2, 3, 5})
		got := removeNthNode(list2, 2)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got != [1, 2, 3, 4, 5]")
			got.Show()
		}
	})
}
