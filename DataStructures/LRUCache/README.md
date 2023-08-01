# Least Recently Used Cache
Is a data structure that stores the `x` most recently used pieces of data to allow for quick "gets". Any data that reaches a "stale" date is removed from the cache when a more fresh entry becomes available.
Any data that is retrieved from the cache has it's "staleness" refreshed.

## Implementation
One implementation of an LRU Cache is to actuall combine two data structures:
1. Double-Linked List
    - This forms the means by which to structure the staleness of the cache and gives quick insert/re-arranging times
2. Hashmap
    - This forms the means by which to store the "location" of the cached data and provides instant look up time.

### Double-Linked List
The head of the double linked list will be the Least Recently Used piece of data and the tail will be the "freshest" piece of data.
- Therefore everytime a piece of data is inserted OR it is accessed through a `Get` function, it will be moved to the tail
    - If a duplicate piece of data (key) is inserted then it is just moved to the tail (and the value can be updated if needed)
- If the capacity of the cache is reached and a new value is inserted, then the Head of the linked list is removed
    - The entry is also removed from the accompanying HashMap

### HashMap
By leveraging a Hashmap we can overcome the "downside" of using just a `Double-Linked List` which is the O(n) look up times.
- When a new entry is made to the cache we can just add it to the HashMap as well
- When a stale entry is removed from the cache we can just delete it from the HashMap
- When a Get() is called we can used the hashmap to immediately index to the node and then move it to the tail
    - And reconnect the nodes on either side

### Code
We will need the following structs for an LRU Cache:
1. Double Linked List Node
    ```Go
    type Node struct {
        Key int
        Val int
        Left *Node
        Right *Node
    }
    ```
2. The LRU Cache which will hold a capacity, a map of nodes and pointers to easily access the head and tail
    ```Go
    type LRUCache struct {
        capacity int
        Head *Node
        Tail *Node
        // The int for the map will be a key, this could be abstracted to be of any type
        cache map[int]*Node
    }
    ```

To construct the LRU Cache we only need user supplied capacity
```Go
    func NewLRUCache(cap int) LRUCache {
        lru := LRUCache {
            capacity: cap,
            Head: &Node{},
            Tail: &Node{},
            cache: make(map[int]*Node),
        }

        // Set the head and tail to point at each other
        lru.Head.Right = lru.Tail
        lru.Tail.Left = lru.Head
    }
```

The `Insert(key, val)` function will need to do three things:
1. Check if the key already exists
    - If it does then update val if necessary and re-allocate to the Tail
2. If it doesn't exist create a new node
    - Insert to the tail
3. Check if the capacity of the cache has been exceeded
    - If it has delete the node nearest the Head
    - Update the map

```Go
func (this *LRUCache) Insert(key, val int) {
    node, ok := this.cache[key]

    // If it exists
    if ok {
        // connect the nodes on either side of it
        node.Left.Right = node.Right
        node.Right.Left = node.Left

        // Move the node to the tail
        oldLeft := this.Tail.Left
        oldLeft.Right = node
        node.Left = oldLeft
        node.Right = this.Tail
        this.Tail.Left = node

        // Update value
        node.Val = val
    } else {
        // If it doesn't exist create a new node and place it at the tail
        node := &Node{Key: key, Val: val}

        // Move the node to the tail
        oldLeft := this.Tail.Left
        oldLeft.Right = node
        node.Left = oldLeft
        node.Right = this.Tail
        this.Tail.Left = node

        // Add to the cache
        this.cache[key] = node
    }

    // Finally check if the capacity has been exceeded
    if len(this.cache) > this.capacity {
        // Grab the node nearest the head
        temp := this.Head.Right
        temp.Left.Right = temp.Right
        temp.Right.Left = temp.Left
        delete(this.cache, temp.key)
    }
}
```

The `Get(key)` function is less complicated and need only concern itself with checking the cache and if it finds a value returning it and moving the node towards the tail

```Go
func (this *LRUCache) Get(key int) int {
    node, ok := this.cache[key]

    if ok {
        node.Left.Right = node.Right
        node.Right.Left = node.Left

        // Move to the tail
        oldLeft := this.Tail.Left
        oldLeft.Right = node
        node.Left = oldLeft
        node.Right = this.Tail
        this.Tail.Left = node
    } else {
        // Or some other appropriate value
        // i.e. could throw an err type
        return -1
    }

    return node.Val
}
```
