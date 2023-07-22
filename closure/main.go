package main

import "fmt"

func closure() func() int {
	value := 0

	return func() int {
		value++
		return value
	}
}

func main() {
	f := closure()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
