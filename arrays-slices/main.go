package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays & Slices")

	var array1 [5]int
	array1[0] = 1
	array1[4] = 5
	fmt.Println(array1)

	array2 := [7]string{"Esse", "é", "um", "texto", "de", "sete", "partes"}
	array2[4] = "com"
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	fmt.Println("Usar um slice garante uma flexibilidade de tamanhos melhor do que a de um array.")

	slice := []int{123, 123}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[len(array2)-1]
	fmt.Println(slice2)
}
