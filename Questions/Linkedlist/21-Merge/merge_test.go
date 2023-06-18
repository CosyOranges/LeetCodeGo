package linkedlist

import (
	linkedlist "dsa/DataStructures/LinkedList"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := &linkedlist.ListNode{}
	list1 = list1.Build([]int{1, 2, 4})
	list2 := &linkedlist.ListNode{}
	list2 = list2.Build([]int{1, 3, 4, 5, 6})

	t.Run("Test empty list", func(t *testing.T) {
		got := mergeTwoLists(list1, nil)

		if !reflect.DeepEqual(got, list1) {
			t.Errorf("Did not return duplicate of list1")
		}
	})

	t.Run("Test empty list", func(t *testing.T) {
		want := &linkedlist.ListNode{}
		want = want.Build([]int{1, 1, 2, 3, 4, 4, 5, 6})
		got := mergeTwoLists(list1, list2)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got:")
			got.Show()
			t.Errorf("want:")
			want.Show()
		}
	})
}
