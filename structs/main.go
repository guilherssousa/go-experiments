package main

import "fmt"

type usuario struct {
	name   string
	age    uint8
	amigos []usuario
}

func main() {
	fmt.Println("Arquivo structs")

	var gui usuario = usuario{
		name: "Guilherme",
		age:  19,
	}

	var yuyu usuario = usuario{
		age:    17,
		name:   "Sayuri",
		amigos: []usuario{gui},
	}

	gael := usuario{"Gael", 20, []usuario{yuyu, gui}}

	fmt.Println(gui, yuyu, gael)
}
