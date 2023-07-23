package main

import "fmt"

func simpleFibonacci(pos int) int {
	if pos <= 1 {
		return pos
	}

	return simpleFibonacci(pos-2) + simpleFibonacci(pos-1)
}

func worker(tasks <-chan int, results chan<- int) {
	for task := range tasks {
		results <- simpleFibonacci(task)
	}
}

func main() {
	buffer_size := 45
	tasks := make(chan int, buffer_size)
	results := make(chan int, buffer_size)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < buffer_size; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 45; i++ {
		result := <-results
		fmt.Println(result)
	}
}
