package linkedlist

import (
	linkedlist "dsa/DataStructures/LinkedList"
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	list1 := &linkedlist.ListNode{}
	list1 = list1.Build([]int{1, 2, 3, 3})
	list2 := &linkedlist.ListNode{}
	list2 = list2.Build([]int{1, 2, 2, 3, 4, 4, 5})

	t.Run("Test 1", func(t *testing.T) {
		want := &linkedlist.ListNode{}
		want = want.Build([]int{1, 2})

		got := deleteDuplicates(list1)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got != want")
			got.Show()
			want.Show()
		}
	})

	t.Run("Test 2", func(t *testing.T) {
		want := &linkedlist.ListNode{}
		want = want.Build([]int{1, 3, 5})

		got := deleteDuplicates(list2)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got != want")
			got.Show()
			want.Show()
		}
	})
}
