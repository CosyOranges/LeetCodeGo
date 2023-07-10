package main

import (
	hashmap "dsa/DataStructures/Hashmap"
	"fmt"
	"strconv"
)

func main() {
	test := hashmap.NewHashmap[string, int](hashmap.Equals[string], hashmap.HashString)
	for i := 0; i < 50; i++ {
		test.Insert(strconv.Itoa(i), i)
	}

	ans, err := test.Get("1")
	fmt.Printf("Test: %v, %v\n", ans, err)

	ans, err = test.Get("2")
	fmt.Printf("Test: %v, %v\n", ans, err)

	ans, err = test.Get("0")
	fmt.Printf("Test: %v, %v\n", ans, err)

	ans, err = test.Get("4")
	fmt.Printf("Test: %v, %v\n", ans, err)
	test.Delete("0")
	ans, err = test.Get("0")
	fmt.Printf("Test: %v, %v\n", ans, err)
	ans, err = test.Get("4")
	fmt.Printf("Test: %v, %v\n", ans, err)
	// ans2, err2 = test.Get("3")
	// fmt.Printf("Test: %v, %v\n", ans2, err2)
}
