# Greedy Algorithms
At it's core a Greedy Algorithm is any algorithm that follows the "problem solving heuristic" of making the locally optimal choice.
- The trade off here is that you may not reach a Global Minimum, but you will reach a minimum that approximates a global one
- Greedy Algorithms are best implemented when speed is the main concern

Some downsides to this approach are:
1. They are short-sighted an unrecoverable
    - This means if a choice is made which does not lead to an optimal solution, there is usually no recovery

2. They can also even produce the worst solution in some cases

## So when are Greedy algorithms appropriate?
-  Mathematically they tend to be the method of choice for Matroids
    - This is a mathematical that generalises the notion of linear independencies from vector spaces to arbitrary sets

- Load balancing
    - This is a problem that lends itself to greedy algorithms as the best way to schedule jobs is to fit the largest job to the machine with the most available space

# Example
A classic example of the Greedy Algorithm problem is the `Distributed Candy` problem:

- You have an integer array of length n
    - Each index represents a child that has been given a rating
    - Your goal is to find the minimum number of candies needed such that:
        - Each child has at least 1 candy
        - Each child has more candy than their neighbour IF the child has a higher rating

- e.g. Children: `[1, 5, 2, 1]`
    - Answer: 7
    - explanation: Candy distributed: `[1, 3, 2, 1]`
