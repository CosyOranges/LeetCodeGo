package binarytree

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	in, post := []int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}

	t.Run("Test PreOrder is as expected", func(t *testing.T) {
		want := []int{3, 9, 20, 15, 7}

		tree := buildTree(in, post)

		got := PreOrderArr(tree)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})
}
