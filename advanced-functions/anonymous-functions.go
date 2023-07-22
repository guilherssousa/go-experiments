package main

import "fmt"

func main() {

	// Like Javacript's IIFE
	func(hello string) {
		fmt.Println("Hello IIFE", hello)
	}("YO DREAM")

	hello := func() {
		fmt.Println("Hello world!")
	}

	hello()
}
