package main

import (
	"fmt"
	"io"
	"os"
)

const POKEMON_RED_BLUE_SAVE_FILE_SIZE = 0x8000

func main() {

	file, err := os.Open("data.bin")

	if err != nil {
		panic(err)
	}

	file_stats, err := file.Stat()

	if err != nil {
		panic(err)
	}

	if file_stats.Size() != POKEMON_RED_BLUE_SAVE_FILE_SIZE {
		panic("file size is not 0x2000")
	}

	defer file.Close()

	buffer := make([]byte, POKEMON_RED_BLUE_SAVE_FILE_SIZE)

	for {
		_, err = file.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	for i := 0; i < len(buffer); i++ {
		if i%16 == 0 {
			fmt.Println()
			fmt.Print("0x", fmt.Sprintf("%04x", i), "\t")
		}
		fmt.Printf("%02X ", buffer[i])
	}
}
