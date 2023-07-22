package main

import (
	"fmt"
	"time"
)

func main() {
	// Concorrência != Paralelismo

	go escrever("Olá, mundo!") // goroutine!
	escrever("Olá, Go!")       //<- Vai fazer o programa acabar
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
