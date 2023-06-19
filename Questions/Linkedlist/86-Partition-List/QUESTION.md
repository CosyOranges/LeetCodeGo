Given the head of a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

The rules are:

- Any number that is less than x has to be before x, and maintain the relative order with thoese that are less than x but already before x.
e.g. `[3,4,1,2], target = 4 -> [3,1,2,4]`, so the order of `[3,1,2]` is maintained.
- Any number that is greater than x but already before x will still be before x, but all of them come after those that are less than x and at the same time maintain their relative order.
e.g. `[3,6,5,4,1,2] target = 4 -> [3,1,2,6,5,4]`
- Any number that is greater than x and after x will only need to maintain their relative order
e.g. `[3,6,5,4,8,1,7,2] target = 4 -> [3,1,2,6,5,4,8,7]`
