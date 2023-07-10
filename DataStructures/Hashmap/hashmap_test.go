package hashmap

import (
	"strconv"
	"testing"
)

func TestHashmap(t *testing.T) {
	testMap := NewHashmap[string, int](Equals[string], HashString)

	// Generate a map that has hash collisions
	for i := 0; i < 50; i++ {
		testMap.Insert(strconv.Itoa(i), i)
	}

	t.Run("Check that keys exist", func(t *testing.T) {
		val, x := testMap.Get("0")

		if !x {
			t.Errorf("Key: %v, not found", val)
		}

		if val != 0 {
			t.Errorf("Expected 1, got: %v", val)
		}
	})

	t.Run("Check delete on known hash collisions", func(t *testing.T) {
		// From the hash function 0 will be inserted at index 33
		// 4 will also be inserted at index 33, therefore if we delete 0
		// We should still be able to find 4
		testMap.Delete("0")
		val, x := testMap.Get("0")

		if x {
			t.Errorf("Key: %v, found when it should be deleted", val)
		}

		val, x = testMap.Get("4")
		if !x {
			t.Errorf("Key: %v, not found", val)
		}
		if val != 4 {
			t.Errorf("Expected 4, got: %v", val)
		}
	})
}
