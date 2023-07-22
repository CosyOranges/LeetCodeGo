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

---

# Code

## Constructing a Segment Tree
When constructing the tree there are two main things to take into consideration
1. Choosing what value to be stored at the node (This is based on the PROBLEM DESCRIPTION)
    - i.e. for a question about interval sums, the sum of the array segment will be stored
2. What should the merge do?
    - i.e. for a range question the child nodes are merged and summed to create the value of the parent node

**Example for the range tree:**
```Go
type SegmentTree struct {
    Nodes []int
    Size int
}

func Constructor(nums []int) *SegmentTree {
    nodes := make([]int, treeSize(len(nums)))

    t := &SegmentTree{Nodes: nodes, Size: len(nums)}
    t.build(nums, 0, 0, len(nums)-1)

    return t
}

func (t *SegmentTree) build(nums []int, node int, left int, right int) {
    if left > right {
        return
    } else if left == right {
        // Then we have divided down to a single element of the input array
        t.Nodes[node] = nums[left]
    } else {
        // We continue to divide
        mid := left + (right - left)/2

        // Build the left subtree
        t.build(nums, 2*node+1, left, mid)

        // Build the right subtree
        t.build(nums, 2*nodes+2, mid+1, right)

        // This merge is going to handle the sum ranges case
        t.Nodes[node] = t.Nodes[2*node+1] + t.Nodes[2*node+2]
    }
}

func treeSize(num int) {
    x := match.Ceil(math.Log2(float64(num)))

    return 2 * int(math.Pow(2, x)) - 1
}
```

### Expand for the Sum Range Case to Update an Element
The algorithm is as follows:
1. If the current node is too far left or right return
2. If left == right then we have found the update index
3. Otherwise we need to traverse further down the tree
    - Search the left tree
    - Search the right tree
4. As we come back up the tree update the node values to account for the updated node
``` Go
func (t *SegmentTree) Update(x, y int) {
    t.update(0, 0, t.Size-1, x, y)
}

func (t *SegmentTree) update(node, left, right, updateIndex, val int) {
    if left > right || left > updateIndex || right < updateIndex {
        // This node is either too far left or too far right
        return
    } else if left == right {
        // We've found the node to update
        t.Nodes[node] = val
    } else {
        mid := left + (right - left)/2

        // Search the left subtree for the index
        t.update(2*node+1, left, mid, updateIndex, val)

        // Search the right subtree for the index
        t.update(2*node+2, left, mid, updateIndex, val)

        // Update the current node
        t.Nodes[node] = t.Nodes[2*node+1] + t.Nodes[2*node+2]
    }
}
```

### Expand to Query for the Sum between Nodes
The algorithm is as follows:
1. If the current index sits inbetween the ranges then we don't need to traverse further down
    - Because the current index is the cummulation of what is below, so return the value
2. If the current index is outside the range then return 0
3. Return the sum of the left and right subtrees
``` Go
func (t *SegmentTree) SumRange(queryL, queryR int) int {
    return t.sumRange(0, 0, t.Size-1, queryL, queryR)
}

func (t *SegmentTree) sumRange(node, left, right, queryL, queryR int) int {
    if queryL <= left && queryR >= right {
        // Found a part of our range
        return t.Nodes[node]
    }

    if queryL > right || queryR < left {
        // Outside of the range
        return 0
    }

    mid := left + (right - left)/2

    return t.sumRange(2*node+1, left, mid, queryL, queryR) + t.sumRange(2*node+2, mid+1, right, queryL, queryR)
}
```
