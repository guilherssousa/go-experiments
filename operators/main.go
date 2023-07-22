package main

import "fmt"

func main() {
	// ARITMÉTICOS

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	var numero int16 = 10
	var numero2 int16 = 25

	soma1 := numero + numero2

	fmt.Println(soma1)

	// FIM DOS ARITMÉTICOS

	// ATRIBuIÇÂO

	var variavel1 string = "sou legal"
	variavel := "STring2"
	fmt.Println(variavel1, variavel)

	// FIM DOS ATRIBUIDORES

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// FIM DOS RELACIONAIS

	// OPERADORES LOGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro, !falso)
	// FIM OPERADORES LOGICOS

	// OPERADOR UNÁRIO
	legal := 10
	legal++
	legal += 4
	fmt.Println(legal)
	legal--
	legal -= 4
	fmt.Println(legal)
	legal %= 2
	fmt.Println(legal)
	// FIM OPERADORES UNÁRIOS

	// OPERADOR TERNÁRIO?
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}
