package linkedlist

import (
	linkedlist "dsa/DataStructures/LinkedList"
	"reflect"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	list1 := &linkedlist.ListNode{}
	list1 = list1.Build([]int{1, 2, 3, 4, 5})

	t.Run("Test standard", func(t *testing.T) {
		want := &linkedlist.ListNode{}
		want = want.Build([]int{1, 4, 3, 2, 5})

		got := reverseBetween(list1, 2, 4)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test right next to each other", func(t *testing.T) {
		want := &linkedlist.ListNode{}
		want = want.Build([]int{1, 3, 4, 2, 5})

		got := reverseBetween(list1, 2, 3)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
