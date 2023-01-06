package main

import (
	"fmt"
	"os"
	"os/exec"
)

// priority list for directories
var DIRECTORY_LIST = []string{"D:/projetos", "C:/Users/Gui/Desktop"}

func main() {
	command := os.Args

	if len(command) < 2 {
		fmt.Println("No folder provided")
		os.Exit(1)
	}

	folder := command[1]

	for _, directory := range DIRECTORY_LIST {
		path := directory + "/" + folder

		info, err := os.Stat(path)

		if err != nil {
			continue
		}

		if info.IsDir() {
			cmd := exec.Command("cmd", "/C", "cd", path)
			cmd.Run()
			os.Exit(0)
		}
	}

	fmt.Println("Folder not found.")
}
