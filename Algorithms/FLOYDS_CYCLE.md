# Floyd's Cycle Finding Algorithm
This algorithm is used to detect a cycle in a linked list object

## Brute Force
A brute force approach to cycle  detection is to create a hashmap that stores the nodes we visit.
- If a `node->next` is already present in the map then we must have a cycle

## Floyd's Algo
The above approach although O(n) in time is also O(n) in space, Floyd's algorithm uses the idea of two runners:
- A slow runner that increments along the list one at a time
- A fast runner that increments along the list two at a time
- With the idea being that if a cycle exists, the fast runner will eventually catch up to the slow runner

```Go
type ListNode struct {
    Val int
    Next *ListNode
}

func CycleDetect(node *ListNode) bool {
    if node == nil {
        return false
    }

    slow, fast := node, node

    for slow != nil && fast != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if fast == slow {
            return true
        }
    }

    return false
}
```
