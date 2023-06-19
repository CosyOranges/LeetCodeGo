package linkedlist

import "fmt"

type LRUCache struct {
	capacity   int
	head, tail *DoubleNode
	cache      map[int]*DoubleNode
}

type DoubleNode struct {
	key  int
	val  int
	Prev *DoubleNode
	Next *DoubleNode
}

func Constructor(capacity int) LRUCache {
	sentinelHead := &DoubleNode{}
	sentinelTail := &DoubleNode{}
	lru := LRUCache{
		capacity: capacity,
		head:     sentinelHead,
		tail:     sentinelTail,
		cache:    make(map[int]*DoubleNode),
	}

	lru.head.Next = lru.tail
	lru.tail.Prev = lru.head

	return lru
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]

	if !ok {
		return -1
	} else {
		// remove node from current position
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev

		// Move node to tail
		oldPrev := this.tail.Prev
		oldPrev.Next = node
		node.Prev = oldPrev
		node.Next = this.tail
		this.tail.Prev = node
	}

	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]

	// If the node exists then move it to the tail of the list
	// We are not adding anything so the list doesn't need to be resized
	if ok {
		// Remove the node by connecting nodes on either side
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev

		// Add the node to the tail
		oldPrev := this.tail.Prev
		oldPrev.Next = node
		node.Prev = oldPrev
		node.Next = this.tail
		this.tail.Prev = node
		node.val = value
	} else {
		// Create the node
		node := &DoubleNode{key: key, val: value}

		// Add the node to the tail
		oldPrev := this.tail.Prev
		oldPrev.Next = node
		node.Prev = oldPrev
		node.Next = this.tail
		this.tail.Prev = node
		node.val = value
		this.cache[key] = node
	}

	// If the list has got too long remove the head
	if len(this.cache) > this.capacity {
		temp := this.head.Next
		temp.Prev.Next = temp.Next
		temp.Next.Prev = temp.Prev
		delete(this.cache, temp.key)
	}
}

func main() {
	// * Your LRUCache object will be instantiated and called as such:
	obj := Constructor(2)
	// param_1 := obj.Get(1)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Printf("Did we find 1: %v\n", obj.Get(1))
	obj.Put(3, 3)
	fmt.Printf("Did we find 2: %v\n", obj.Get(2))
	fmt.Printf("Did we find 1: %v\n", obj.Get(1))
}
