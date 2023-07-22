package main

import "fmt"

var n int

// Isso aqui roda antes da função main.
func init() {
	fmt.Println("Ordem 1")
	n = 5
}

func main() {
	fmt.Println("Ordem 2")
	fmt.Println(n)
}
