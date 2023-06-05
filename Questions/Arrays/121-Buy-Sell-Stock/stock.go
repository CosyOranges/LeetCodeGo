package arrays

import "fmt"

func maxProfit(prices []int) int {
	profit := 0

	// brute force algorithm O(nlogn)
	// for i := 0; i < len(prices); i++ {
	// 	for j := i; j < len(prices); j++ {
	// 		k := prices[j] - prices[i]
	// 		if k > profit {
	// 			profit = k
	// 		}
	// 	}
	// }

	// An O(n) solution is to track the minimum stock price
	minPrice := prices[0]

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}

	return profit
}

func maxProfitII(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}

func main() {
	stock := []int{7, 1, 5, 3, 6, 2}

	ans := maxProfit(stock)

	fmt.Printf("Answer Part1: %v\n", ans)

	ans = maxProfitII(stock)

	fmt.Printf("Answer Part2: %v\n", ans)
}
