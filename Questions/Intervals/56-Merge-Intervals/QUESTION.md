Given an array of intervals:
- Where `intervals[i] = [start_i, end_i]`
- Merge all **overlapping** intervals and return an array of non overlapping intervals
    - This must cover all the intervals of the input array

---

Example 1:
```
Input = [[1, 3], [2, 6], [8, 10], [15, 18]]

Output = [[1, 6], [8, 10], [15, 18]]
```

Example 2:
```
Input = [[1, 3], [2, 6], [6, 9]]

Output = [[1, 9]]
```
