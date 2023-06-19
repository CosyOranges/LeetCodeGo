Given an array of positive integers nums and a positive integer target, return the minimal length of a 
subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

# Solution

## `O(n^2)`
The `O(n^2)` solution would involve an adjustable sliding window that starts by looking for the ideal situation which is that there exists one index in the array >= target
- If not then increase window size by one and search again

```Go
func minSubArray(nums []int, target int) int {
    sum, window := 0, 1

    for window < len(nums)+1 {
        start, end := 0, window
        for end < len(nums)+1; end++ {
            for i := start; i < end; i++ {
                sum += nums[i]
            }
            if sum >= target {
                return window
            }
            sum = 0
            end++
            start++
        }
        window++
    }

    return 0
}
```

## O(n)
The correct solution for this is to use two points left and right
- Both start at index 0
- We increment right and sum the numbers until sum >= target
    - Once we have summed enough if the length of the summed numbers is less than the minLength
        - Replace the min length
    - Then subtract the value of the left index from the sum
    - increment the left index up the array
- If the sum is no longer >= target increment the right index again

```Go
func minSubArray(nums []int, target int) int {
    l, sum, minLength := 0, 0, len(nums) + 1
    
    for r := 0; r < len(nums); r++ {
        sum += nums[r]
        for sum >= target {
            if minLength > (r - l) + 1 {
                minLength = (r - l) + 1
            }
            sum -= nums[l]
            l++
        }
    }

    // We did not find a subarray that satisfied the target
    if minLength == len(nums) + 1 {
        return 0
    }

    return minLength
}
```
