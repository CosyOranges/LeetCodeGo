Question: Given a list of inorder and postorder keys for a binary tree
    construct the tree and then print a DFS of it
#
Notes: 
     - Not a Binary Search Tree
     - No duplicate keys
Orders if they do not know

- Inorder (Left, Root, Right)

- Preorder (Root, Left, Right) 

- Postorder (Left, Right, Root) 
#
Input : 

`in[]   = {2, 1, 3}`

`post[] = {2, 3, 1}`

Output : Root of below tree
```
     1
   /   \
  2     3 
```

STDOUT: 1 2 3
#
Input : 

`in[]   = {4, 8, 2, 5, 1, 6, 3, 7}`

`post[] = {8, 4, 5, 2, 6, 7, 3, 1} `

Output : Root of below tree
```
         1
      /     \
    2        3
  /    \   /   \
 4     5   6    7
   \
     8
```
STDOUT: 1 2 4 8 5 3 6 7

Final Output: DFS (preorder) of tree