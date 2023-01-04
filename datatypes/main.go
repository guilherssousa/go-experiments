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

	// int32 = rune
	// int8 = byte

	var number2 rune = 123456
	fmt.Println(number2)

	var number3 byte = 123
	fmt.Println(number3)

	// we have 2 types of float
	// float32 and float64

	var number4 float32 = 123.45
	fmt.Println(number4)

	var number5 float64 = 12323232.45
	fmt.Println(number5)

	// strings
	// string is a sequence of bytes

	var str string = "Uma string"
	fmt.Println(str)

	// golang does not have char
	// putting a value between single quotes
	// will be a char in ASCII

	char := '%'
	fmt.Println(char)

	// boolean

	var boolean bool = true
	fmt.Println(boolean)

	// error
	// error is a interface
	// type error interface {
	// 	Error() string
	// }

	// zero values

	var number6 int
	fmt.Println(number6)

	var number7 float64
	fmt.Println(number7)

	var str2 string
	fmt.Println(str2)

	var boolean2 bool
	fmt.Println(boolean2)
}