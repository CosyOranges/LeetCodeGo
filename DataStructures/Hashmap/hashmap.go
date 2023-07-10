package hashmap

import (
	"fmt"

	"github.com/segmentio/fasthash/fnv1a"
)

const MAP_SIZE = 50

type HashmapNode[K any, V any] struct {
	Key  K
	Val  V
	Next *HashmapNode[K, V]
}

type ops[T any] struct {
	equals func(a, b T) bool
	hash   func(t T) uint64
}

type Hashmap[K any, V any] struct {
	data []*HashmapNode[K, V]

	size int
	ops  ops[K]
}

/*
HashFn is a function that returns the hash of it's input
*/
type HashFn[T any] func(t T) uint64

/*
EqualsFn is a function that is used to compare generic types
Go doesn't have generics so this implements a wrapper around the '==' operator
*/
type EqualsFn[T any] func(a, b T) bool

/*
Return a new hashmap of any type.

- Key input is the hash function this has to match the type of the key for the hasmap you are declaring
*/
func NewHashmap[K, V any](equals EqualsFn[K], hash HashFn[K]) *Hashmap[K, V] {
	return &Hashmap[K, V]{
		data: make([]*HashmapNode[K, V], MAP_SIZE),
		size: MAP_SIZE,
		ops: ops[K]{
			equals: equals,
			hash:   hash,
		},
	}
}

/*
Insert a new key value pair into the hashmap

Behind the scenes this will check whether the current linked list of HashmapNodes
already contains the key, if it does the value will be overwritten otherwise the
new HashmapNode will be appended to the current list
  - This has the potential to grow linear linked lists
*/
func (m *Hashmap[K, V]) Insert(key K, value V) {
	index := m.getIndex(key)
	fmt.Printf("Insert Index: %v\n", index)

	if m.data[index] == nil {
		m.data[index] = &HashmapNode[K, V]{Key: key, Val: value}
		return
	} else {
		curr := m.data[index]
		for curr.Next != nil {
			// If we encounter the key replace the value
			if m.ops.equals(key, curr.Key) {
				curr.Val = value
				return
			}
			curr = curr.Next
		}
		curr.Next = &HashmapNode[K, V]{Key: key, Val: value}
	}

	return
}

/*
Retrieve a value if the key exists in the map, otherwise return false
*/
func (m *Hashmap[K, V]) Get(key K) (V, bool) {
	index := m.getIndex(key)

	if m.data[index] != nil {
		curr := m.data[index]
		for ; curr != nil; curr = curr.Next {
			if m.ops.equals(curr.Key, key) {
				return curr.Val, true
			}
		}
	}

	var v V
	return v, false
}

/*
Delete a key and value from the map if it's found
*/
func (m *Hashmap[K, V]) Delete(key K) {
	index := m.getIndex(key)

	head := m.data[index]
	pos := 0

	if m.data[index] != nil {
		if pos == 0 && m.ops.equals(m.data[index].Key, key) {
			fmt.Printf("Here at position2: %v\n", pos)

			m.data[index] = m.data[index].Next
			return
		}
		prev, curr := head, head.Next
		for curr != nil {
			if m.ops.equals(curr.Key, key) {
				fmt.Printf("Here at position: %v\n", pos)
				temp := curr
				prev.Next = temp.Next
				return
			}
			prev = curr
			curr = curr.Next
			pos++
		}

	}
	m.data[index] = head

	return
}

/*
Hash an input we assume here that the input is converted to some form of a string.
  - We are using the Jenkins one-at-a-time hashing algorithm for no other reason other
    than it's easy to implement
*/
func hash(key uint64) (hash uint64) {
	hash = 0
	key ^= key >> 33
	key *= 0xff51afd7ed558ccd
	key ^= key >> 33
	key *= 0xc4ceb9fe1a85ec53
	key ^= key >> 33
	return key
}

/*
Wrap the '==' operator for comparable types
*/
func Equals[T comparable](a, b T) bool {
	return a == b
}

/*
The getIndex will modulo the hash output to ensure the key falls within the range
of the map size.
*/
func (m *Hashmap[K, V]) getIndex(key K) int {
	return int(m.ops.hash(key)) & (m.size - 1)
}

func (m *Hashmap[K, V]) HashUint64(u uint64) uint64 {
	return hash(u)
}
func (m *Hashmap[K, V]) HashUint32(u uint32) uint64 {
	return hash(uint64(u))
}
func HashUint16(u uint16) uint64 {
	return hash(uint64(u))
}
func HashUint8(u uint8) uint64 {
	return hash(uint64(u))
}
func HashInt64(i int64) uint64 {
	return hash(uint64(i))
}
func HashInt32(i int32) uint64 {
	return hash(uint64(i))
}
func HashInt16(i int16) uint64 {
	return hash(uint64(i))
}
func HashInt8(i int8) uint64 {
	return hash(uint64(i))
}
func HashInt(i int) uint64 {
	return hash(uint64(i))
}
func HashUint(i uint) uint64 {
	return hash(uint64(i))
}
func HashString(s string) uint64 {
	return fnv1a.HashString64(s)
}
func HashBytes(b []byte) uint64 {
	return fnv1a.HashBytes64(b)
}
