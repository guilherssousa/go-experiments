package main

import "fmt"

func main() {
	var texto string = "Um texto muito legal"
	var numero int = 10

	// inferindo o tipo baseado do valor
	implicita := "Uma string implicita"

	fmt.Println(texto, numero, implicita)

	var (
		legal uint8 = 255
		legal2 uint16 = 65535
	)

	fmt.Println(legal, legal2)

	const constante = "Uma constante"

	// infer types to constant
	const constante2 = "Uma constante 2"

	fmt.Println(constante, constante2)
}