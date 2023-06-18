package linkedlist

import "testing"

func TestHasCycleFloyd(t *testing.T) {
	linkedList := &ListNode{}
	linkedList = linkedList.Build([]int{3, 2, 0, -4})

	t.Run("Test no cylce", func(t *testing.T) {
		if hasCycle(linkedList) {
			t.Errorf("got: true, want: false")
		}
	})

	t.Run("Test with cylce", func(t *testing.T) {
		linkedList.CreateCycle(2, 4)
		if !hasCycle(linkedList) {
			t.Errorf("got: true, want: false")
		}
	})
}
