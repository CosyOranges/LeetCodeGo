package arrays

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	set map[int]bool
}

func Constructor() RandomizedSet {
	return RandomizedSet{set: make(map[int]bool)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; !ok {
		this.set[val] = true
		return true
	}

	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if this.set[val] {
		delete(this.set, val)
		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	k, count := len(this.set), 0

	if k == 0 {
		return -1
	}

	random_index := rand.Intn(k)

	for key := range this.set {
		if count == random_index {
			return key
		}
		count++
	}

	return -1
}

func main() {
	obj := Constructor()
	param_1 := obj.Insert(1)
	fmt.Printf("Insert(1): %v\n", param_1)
	param_2 := obj.Remove(2)
	fmt.Printf("Remove(2): %v\n", param_2)

	obj.Insert(2)
	obj.Insert(3)
	obj.Insert(4)
	obj.Insert(5)
	obj.Insert(6)
	obj.Insert(7)
	param_3 := obj.Remove(2)
	fmt.Printf("Remove(2): %v\n", param_3)
	param_4 := obj.GetRandom()
	fmt.Printf("Remove(2): %v\n", param_4)
}
