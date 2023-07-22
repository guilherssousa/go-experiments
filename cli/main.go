package main

import (
	"cli/app"
	"log"
	"os"
)

func main() {
	app := app.Generate()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
