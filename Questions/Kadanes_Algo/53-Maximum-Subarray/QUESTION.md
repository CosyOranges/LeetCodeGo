Given an integer array nums, find the subarray with the largest sum, and return its sum.

# Solution
The solution to such a question makes use of Kadanes algorithm where we track a global and local sum

- If the current index is ever greater than the global local sum + current index then the local sum is reset to the current index
    - Otherwise the local sum is current index + local sum
- If local sum is greater than global sum, replace

```Go
func maxSubArray(nums []int) int {
    globalSum, localSum := nums[0], nums[0]

    for i := 1; i < len(nums); i++ {
        if nums[i] > localSum + nums[i] {
            localSum = nums[i]
        }

        if localSum > globalSum {
            globalSum = localSum
        }
    }
    return globalSum
}
```