package main

import "fmt"

func soma(numeros ...int) (soma int) {
	soma = 0
	for _, i := range numeros {
		soma += i
	}
	return
}

func main() {
	total := soma(123, 123, 123, 123, 123, 123, 113)
	fmt.Println(total)
}
