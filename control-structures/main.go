package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 15

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := 15; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	}

	// 👇 Isso aqui é tipo o With do Python :)
	if testeBoolean := true; testeBoolean {
		fmt.Println("Variável criada como true")
	} else {
		fmt.Println("Variável criada como false")
	}
}
