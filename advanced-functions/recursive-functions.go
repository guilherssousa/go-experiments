package main

import (
	"fmt"
	"time"
)

// Implement cache so Fibonacci won't explode my memory and stack <3
var cache = map[uint]uint{1: 1, 2: 1, 3: 2}
var queries = 0

func fibonacci(position uint) (value uint) {
	value = 0

	if position <= 1 {
		return position
	}

	if cache[position] != 0 {
		queries++
		return cache[position]
	}

	result := fibonacci(position-2) + fibonacci(position-1)

	cache[position] = result

	return result
}

func main() {
	start := time.Now()
	position := 500
	value := fibonacci(uint(position))

	fmt.Println("Position:", position)
	fmt.Println("Value:", value)
	fmt.Println("Cached sequences:", len(cache))
	fmt.Println("Cache hit:", queries)
	fmt.Println("Time elapsed:", time.Now().Sub(start))
}
