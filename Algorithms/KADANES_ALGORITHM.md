# Kadane's Algorithm
This algorithm is typically used to find the maximum contiguous sub-array in an array.
 - For example: [1, 2, 3, -4, -5]
    - The maximum sub-array is [1, 2, 3] = 6

## Assumptions
There are a few assumptions made
1. There are NO empty sub-arrays within the array
2. `globalMaxSum` denotes the maximum sum of the sub-array
3. `localMaxSum` at index `i` denotes the maximum sum of the sub-array starting with `input[i]`
4. `n` is the size of `input`

### Brute Force Algorithm
The head-on way to tackle this problem is to:
1. loop 0 -> n
    - then have an inner loop from current index -> n
    - calculate the sum of current index -> n
    - if greater than trackedMax replace
2. continue

```Go
globalMaxSum := input[0]
for i := 0; i <len(input); i++ {
    localMaxSum := math.MinInt
    for j := i; j < len(input); j++ {
        localMaxSum += input[j]
    }

    if localMaxSum > globalMaxSum {
        globalMaxSum = localMaxSum
    }
}

return globalMaxSum
```

### Kadane's Approach
Instead of looping through every possible sub-array in the `input`, Kadane's algorithm says that:
- The `localMaxSum` is the maximum of `input[i]` or `input[i] + localMaxSum`
- If this ever exceeds the `globalMaxSum` then that can be updated

```Go
globalMaxSum := math.MinInt
localMaxSum := math.MinInt

for i := 0; i < len(input); i++ {
    if input[i] > localMaxSum + input[i] {
        localMaxSum = input[i]
    } else {
        localMaxSum += input[i]
    }

    if localMaxSum > globalMaxSum {
        globalMaxSum = localMaxSum
    }
}

return globalMaxSum
```
