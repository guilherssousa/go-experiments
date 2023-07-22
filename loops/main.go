package main

import (
	"fmt"
)

func main() {
	lyrics := [...]string{
		"Yaballa Bahiya",
		"Yaballa Bahiya",
		"Let's cast a spell!",
		"Oh yeah",
		"Yaballa Bahiya moha ima moha irura",
	}

	single_line := "I got you under my skin"

	for _, l := range lyrics {
		fmt.Println(l)
	}

	for _, l := range single_line {
		fmt.Println(string(l))
	}

	for {
		fmt.Println("Executando infinitamente")
	}
}
