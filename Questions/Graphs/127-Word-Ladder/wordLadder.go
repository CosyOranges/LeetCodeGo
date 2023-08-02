package graphs

import "fmt"

// Helper node struct
type Node struct {
	Word string
	Len  int
}

func wordLadder(beginWord, endWord string, wordList []string) int {
	/*
		The solution to this problem is likely a Breadth First Search
		The improvement upon a standard BFS solution is to use a binary BFS (or a meet in the middle)
		- This is where you use two queues, one that starts from the source, and another that starts from the end
		- They then will meet at some point (if a solution exists) and the solution can be found by summing the two (-1 because they share a node)
	*/
	// Check if endWord exists in wordList
	var found bool = false
	for _, x := range wordList {
		if x == endWord {
			found = true
		}
	}
	if !found {
		return 0
	}

	// Initialise both a queue and map for the start and end
	queueStart, queueEnd := []*Node{}, []*Node{}
	mapStart, mapEnd := map[string]int{}, map[string]int{}

	// Add the start word to our start queue and map
	queueStart = append(queueStart, &Node{Word: beginWord, Len: 1})
	mapStart[beginWord] = 1

	// Add the end word to our end queue and map
	queueEnd = append(queueEnd, &Node{Word: endWord, Len: 1})
	mapEnd[endWord] = 1

	for len(queueStart) > 0 && len(queueEnd) > 0 {
		// Get current start word
		var word1 *Node = queueStart[0]
		queueStart = queueStart[1:]

		// Get current end word
		var word2 *Node = queueEnd[0]
		queueEnd = queueEnd[1:]
		fmt.Printf("Start: %v | End: %v\n", word1, word2)
		// Check all the neighbours of the first word
		for _, x := range wordList {
			// Check if the word is "adjacent" (i.e. differs by only 1 char) and not already used
			_, ok := mapStart[x]
			if isAdjacent(word1.Word, x) && !ok {
				var temp *Node = &Node{Word: x, Len: word1.Len + 1}
				queueStart = append(queueStart, temp)
				mapStart[x] = word1.Len + 1

				// Check if we've found the end node
				if temp.Word == endWord {
					return temp.Len
				}

				// Check if start and end have met
				if _, ok := mapEnd[temp.Word]; ok {
					return temp.Len + mapEnd[temp.Word] - 1
				}
			}
		}

		// Check all the neighbours of the second word
		for _, x := range wordList {
			// Check if the word is adjacent
			_, ok := mapEnd[x]
			if isAdjacent(word2.Word, x) && !ok {
				var temp *Node = &Node{Word: x, Len: word2.Len + 1}
				queueEnd = append(queueEnd, temp)
				mapEnd[x] = word2.Len + 1

				// Check if we've found the end node
				if temp.Word == beginWord {
					return temp.Len
				}

				// Check if the start and end have met
				if _, ok := mapStart[temp.Word]; ok {
					return temp.Len + mapStart[temp.Word] - 1
				}
			}
		}
	}

	// Otherwise no solution was found
	return 0
}

func isAdjacent(x, y string) bool {
	var count int = 0

	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			count++
		}
	}

	if count > 1 {
		return false
	}

	return true
}

func main() {
	fmt.Printf("Solution: %v\n", wordLadder("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}
