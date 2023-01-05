package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

const POKEMON_RED_BLUE_SAVE_FILE_SIZE = 0x8000

func open_save_file(path string) ([]byte, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	file_stats, err := file.Stat()

	if err != nil {
		return nil, err
	}

	if file_stats.Size() != POKEMON_RED_BLUE_SAVE_FILE_SIZE {
		return nil, errors.New("file size is not 0x2000")
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
	return buffer, nil
}

func main() {
	save, err := open_save_file("data.bin")

	if err != nil {
		panic(err)
	}

	fmt.Println(save)
}
