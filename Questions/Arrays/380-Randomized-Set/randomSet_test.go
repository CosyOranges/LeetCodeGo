package arrays

import "testing"

func TestRandomSet(t *testing.T) {
	obj := Constructor()
	obj1 := Constructor()
	t.Run("Test insertion", func(t *testing.T) {
		got := obj.Insert(1)
		if !got {
			t.Errorf("Attempted to insert 1 when none existed, got: %v", got)
		}

		got = obj.Insert(1)
		if got {
			t.Error("Successfully inserted 1 when it already exists... shouldn't be possible")
		}
	})

	t.Run("Test deletion", func(t *testing.T) {
		got := obj.Remove(1)
		if !got {
			t.Errorf("Attempted to delete 1 when it existed, got: %v", got)
		}

		got = obj.Remove(1)
		if got {
			t.Error("Successfully deleted 1 when it doesn't exist... shouldn't be possible")
		}
	})

	t.Run("Test get random on empty set", func(t *testing.T) {
		got := obj1.GetRandom()
		want := -1

		if got != -1 {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test get random", func(t *testing.T) {
		obj1.Insert(2)
		obj1.Insert(3)
		obj1.Insert(4)
		obj1.Insert(5)
		obj1.Insert(6)
		obj1.Insert(7)

		got := obj1.GetRandom()
		if got < 2 {
			t.Errorf("got: %v is less than minimum value in set", got)
		}
	})
}
