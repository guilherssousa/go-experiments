package main

import "fmt"

func recuperateExecution() {
	fmt.Println("Tentativa de recuperar a execução")

	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func isStudentApproved(n1, n2 float64) bool {
	defer recuperateExecution()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média não pode ser exatamente seis.")
}

func main() {
	fmt.Println(isStudentApproved(6, 6))
}
