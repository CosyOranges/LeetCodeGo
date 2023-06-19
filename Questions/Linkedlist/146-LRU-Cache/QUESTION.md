Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:

LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
int get(int key) Return the value of the key if the key exists, otherwise return -1.
void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
The functions get and put must each run in O(1) average time complexity.

# Solution

The solution is to combine a hashmap and a doubly linked list.
- The hashmap allows us to find the node in O(1) time
- The doubly linked list allows us to delete nodes in O(1) time

### The Hashmap
This will be of the form: `map[int]*Node`
- Where `int` is the `key` of the node

### The Node
This will be of the form:
```Go
type Node struc {
    key     int
    val     int
    next    *Node
    prev    *Node
}
```

### The LRU Cache
- This will have two "sentinel" nodes which don't hold any key, values
- These aare only there to make insertion and deletion simpler

### Constructor
```Go
type LRUCache struct {
    capacity    int
    cache       map[int]*Node
    head        *Node
    tail        *Node
}
```
The `Constructor(capacity)` will do 3 things:
1. Create a sentinel Head node
2. Create a sentinal Tail node
3. Assign the `Head.next` -> Tail
4. Assign the `Tail.prev` -> Head
5. make the map[int]*Node
6. Assign the capacity and return the LRUCache

### Put Function
The `Put(key, val)` will do 4 things:
1. Check if the node exists using the key
2. If the node does exist
    - Remove it from the head and place it at the tail
    - Update the value of the node with the new value
3. If the node does not exist
    - Create it at the tail
    - Point the map[key] at this node
4. If the map is now longer than the capacity of the cache remove the head

### Get Function
The `Get(key)` will do 2 things:
1. Check if the key exists
    - If not return -1
2. If it does move the node to the tail
    - Return nodes value
