package arrays

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	ans, startIndex, globalGas, currentGas := -1, 0, 0, 0

	for i := 0; i < len(gas); i++ {
		globalGas += gas[i] - cost[i]
		currentGas += gas[i] - cost[i]
		if currentGas < 0 {
			startIndex = i + 1
			currentGas = 0
		}
	}

	if globalGas < 0 {
		return ans
	}
	return startIndex
}

func main() {
	gas, cost := []int{3, 1, 1}, []int{1, 2, 2}
	fmt.Printf("ANS: %v\n", canCompleteCircuit(gas, cost))

	gas, cost = []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}
	fmt.Printf("ANS: %v\n", canCompleteCircuit(gas, cost))
}
