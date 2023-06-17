package hashmap

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	yes := []int{1, 2, 3, 1, 2, 3}

	t.Run("Test no case", func(t *testing.T) {
		if containsNearbyDuplicate(yes, 2) {
			t.Errorf("got: true, want: false")
		}
	})

	t.Run("Test yes case", func(t *testing.T) {
		if !containsNearbyDuplicate(yes, 3) {
			t.Errorf("got: false, want: true")
		}
	})
}
