package main

import "fmt"

func main() {
	buffer_size := 2
	// O segundo argumento indica que tem um buffer!
	// A diferença é que só vai bloquear quando
	// atingir a capacidade máxima
	// Um canal normal sempre vai bloquear.
	canal := make(chan string, buffer_size)
	canal <- "Olá mundo!"
	canal <- "Codando em Go..."

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
