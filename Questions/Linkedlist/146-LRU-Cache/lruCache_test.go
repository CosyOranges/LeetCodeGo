package linkedlist

import "testing"

func TestLRUCache(t *testing.T) {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)

	t.Run("Get valid key", func(t *testing.T) {
		want := 1

		got := obj.Get(1)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test 2 removed from cache", func(t *testing.T) {
		want := -1
		obj.Put(3, 3)
		got := obj.Get(2)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
