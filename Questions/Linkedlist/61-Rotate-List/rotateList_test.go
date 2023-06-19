package linkedlist

import (
	linkedlist "dsa/DataStructures/LinkedList"
	"reflect"
	"testing"
)

func TestRotateRight(t *testing.T) {
	list := &linkedlist.ListNode{}
	list = list.Build([]int{0, 1, 2, 3, 4, 5})

	t.Run("Test rotate k < length(nums)", func(t *testing.T) {
		want := &linkedlist.ListNode{}
		want = want.Build([]int{4, 5, 0, 1, 2, 3})

		got := rotateRight(list, 2)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got was not rotated by two.")
			got.Show()
			want.Show()
		}
	})

	t.Run("Test rotate k > length(nums)", func(t *testing.T) {
		list = list.Build([]int{0, 1, 2, 3, 4, 5})
		want := &linkedlist.ListNode{}
		want = want.Build([]int{5, 0, 1, 2, 3, 4})

		got := rotateRight(list, 7)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got was not rotated by two.")
			got.Show()
			want.Show()
		}
	})
}
