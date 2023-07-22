package main

import (
	"fmt"
	"time"
)

// Concorrência != Paralelismo
func main() {
	channel := make(chan string)
	go escrever("Hello, world!", channel)

	for message := range channel {
		fmt.Println(message)
	}
}

func escrever(texto string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- texto
		time.Sleep(time.Second)
	}

	close(channel)
}
