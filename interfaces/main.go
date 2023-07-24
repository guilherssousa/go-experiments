package main

import (
	"fmt"
	"nice-testing/shapes"
)

func escreverArea(f shapes.Forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.Area())
}

func main() {
	r := shapes.Retangulo{
		Altura:  10,
		Largura: 15,
	}
	escreverArea(r)
	c := shapes.Circulo{Raio: 10}
	escreverArea(c)
}
