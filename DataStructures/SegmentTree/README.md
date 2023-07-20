# What is a Segment Tree?
It is a data structure that stores information about a range of elements in its nodes.
Each level of the array stores information about a different subset of the array.

---

## Structure of the Tree
The Segment Tree works ont the principle of divide an conquer.
- At each level, we divide the array segments into two parts.
    - If the array has `N` elements `[0, ..., N-1]` Then it will be divided into:
        - `[0, ..., (N/2)-1]`
        - `[(N/2), ..., N-1]`

- It is then recrusively split until there is an individual element
- Ultimately it looks like a binary tree however, there is not necessarily any order to the nodes (except for that within the initial array)

## Constructing a Segment Tree
When constructing the tree there are two main things to take into consideration
1. Choosing what value to be stored at the node (This is based on the PROBLEM DESCRIPTION)
    - i.e. for a question about interval sums, the sum of the array segment will be stored
2. What should the merge do?
    - i.e. for a range question the child nodes are merged and summed to create the value of the parent node

**Example for the range tree:**
```Go
type Node struct {
    Val int
    Left *Node
    Right *Node
}

type SegmentTree struct {
    Root *Node
}

func Constructor(nums []int) *SegmentTree {
    if len(nums) == 0 {
        return &SegmentTree{Root: nil}
    }

    
}
```
