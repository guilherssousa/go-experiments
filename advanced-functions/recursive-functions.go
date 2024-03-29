package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// Implement cache so Fibonacci won't explode my memory and stack <3
var cache = map[uint]uint{1: 1, 2: 1, 3: 2}
var queries = 0
var calculus = 0

func fibonacci(position uint) (value uint) {
	value = 0
	calculus++

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

	position, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	start(uint(position))
}

func start(position uint) {
	if position >= 191_600_00 {
		log.Fatal("Please use a input less than 191_600_00.")
	}

	start := time.Now()

	value := fibonacci(position)

	fmt.Println("Position:", position)
	fmt.Println("Value:", value)
	fmt.Println("Cached sequences:", len(cache))
	fmt.Println("Cache hit:", queries)
	fmt.Println("Ran functions:", calculus)
	fmt.Println("Time elapsed:", time.Now().Sub(start))
}
