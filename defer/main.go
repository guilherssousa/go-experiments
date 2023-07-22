package main

import "fmt"

func func1() {
	fmt.Println("Executando a função 1")
}

func func2() {
	fmt.Println("Executando a função 2")
}

func isStudentApproved(n1, n2 float32) bool {
	defer fmt.Println("Média calculada.")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	fmt.Println("Calculando média...")

	media := (n1 + n2) / 2

	return media >= 6

}

func main() {
	defer func1()
	func2()

	fmt.Println(isStudentApproved(5.5, 6.6))
}
