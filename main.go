package main

import (
	"fmt"
	"math/rand"
	"time"
)

var list = make([]int, 10)

func main() {
	for i := 0; i < 9; i++ {
		list[i] = rand.Intn(10)
	}
	fmt.Println("List created: ", list)
	fmt.Println("Because the Big O notation of this algorithm is O(n!), we should be expecting this to take 39,916,800 (39.916 million) iterations to complete on average")
	iteration := 0
	start := time.Now()
	for {
		iteration++
		rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
		if check() {
			break
		}

	}

	fmt.Println("Time taken (ms): ", time.Since(start))
	fmt.Println("Sorted List: ", list)
	fmt.Println("Iterations: ", iteration)
}

func check() bool {
	var lastNum = 0
	for i := 0; i < 10; i++ {
		if lastNum > list[i] {
			return false
		} else {
			lastNum = list[i]
		}
	}
	return true
}
