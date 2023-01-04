package main

import "fmt"

func main() {
	// learning about the golang datatypes
	// int8, int16, int32, int64

	var number int16 = 100

	// when you only use int, it will use the
	// computer architecture to define the int size.

	// int8 = -128 to 127
	// int16 = -32768 to 32767
	// int32 = -2147483648 to 2147483647
	// int64 = -9223372036854775808 to 9223372036854775807

	// int will be equivalent to int32 or int64
	// depending on the computer architecture

	fmt.Println(number)
}