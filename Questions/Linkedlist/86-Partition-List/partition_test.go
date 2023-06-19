package linkedlist

import (
	linkedlist "dsa/DataStructures/LinkedList"
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	list := &linkedlist.ListNode{}
	list = list.Build([]int{1, 4, 3, 2, 5, 2})

	t.Run("Test normal case", func(t *testing.T) {
		want := &linkedlist.ListNode{}
		want = want.Build([]int{1, 2, 2, 4, 3, 5})

		got := partition(list, 3)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got was not partitioned correctly")
			got.Show()
			want.Show()
		}
	})

	t.Run("Test with x not in list", func(t *testing.T) {
		want := &linkedlist.ListNode{}
		want = want.Build([]int{1, 2, 2, 4, 3, 5})

		got := partition(list, 9)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got was not partitioned correctly")
			got.Show()
			want.Show()
		}
	})
}
