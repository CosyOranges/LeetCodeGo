package arrays

import "fmt"

func candy(ratings []int) int {
	// Solve this with the brutal algorithm first:
	/*
			1. Outer while loop to continue until flag is satisfied
			2. Inner loop we loop through the array and increment each child if needed
				a. Here we check three things
					1. If the child has been assigned any candies at all, if not make it = 1
					2. If the child has been assigned more candies than it's right neighbour (if needed)
						- Taking into account that i must be < n - 1
						- If not the Greedy solution is just have the current child have 1 more than it's nieghbour
						- i.e. candies[i] = candies[i+1]+1
					3. If the child has been assigned more candies than it's left neighbour (if needed)
						- Taking into account that i must be > 0
						- If not the Greedy solution is just have the current child have 1 more than it's nieghbour
						- i.e. candies[i] = candies[i-1]+1

		candies := make([]int, len(ratings))
		solved := true
		for solved {
			solved = false
			for i := 0; i < len(candies); i++ {
				if candies[i] == 0 {
					candies[i]++
				}

				if i != len(candies)-1 && ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
					candies[i] = candies[i+1] + 1
					solved = true
				}

				if i > 0 && ratings[i] > ratings[i-1] && candies[i] <= candies[i-1] {
					candies[i] = candies[i-1] + 1
					solved = true
				}
			}
			fmt.Printf("Candies: %v\n", candies)
		}

		ans := 0
		for _, x := range candies {
			ans += x
		}
	*/

	/*
		The more optimal solution is to apply a greedy algorithm:
		Here the problem can be divided into two because each child has two neighbours
		but we can simplify this by only considering one set of neighbours at a time
		1. We create a Left array that is only ever concerned with the left neighbour of the child
			- if ratings[i] >= ratings[i-1] then candies[i] = candies[i-1]+1
		2. We create a Right array that is only ever concerned with the right neighbour of the child
			- if ratings [i] >= ratings[i+1] then candies[i] = candies[i+1]+1
		3.
	*/
	left, right := make([]int, len(ratings)), make([]int, len(ratings))

	// Set everything to 1
	for i := 0; i < len(left); i++ {
		left[i], right[i] = 1, 1
	}
	for i := 0; i < len(left); i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	for i := len(right) - 1; i >= 0; i-- {
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	// Now we can find our answer by summing the maximum values at each index from left and right
	ans := 0
	for i := 0; i < len(left); i++ {
		if left[i] >= right[i] {
			ans += left[i]
		} else {
			ans += right[i]
		}
	}

	return ans
}

func main() {
	ratings := []int{1, 2, 3, 3, 3, 2, 1}

	fmt.Printf("Ans: %v\n", candy(ratings))
}
