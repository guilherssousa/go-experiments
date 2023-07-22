package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func() string {
		fmt.Println("função f")
		return "hello world"
	}

	f()

  resultadoSoma, _ := calculosMatematicos(125, 125)
  fmt.Println(resultadoSoma)
}
