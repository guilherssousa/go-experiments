package main

import (
	"aoc/utils"
	"bufio"
	"fmt"
	"strings"
)

func main() {
  input := utils.ReadInputs("inputs/day-3.txt")
  defer input.Close()

  l := 0
  sc := bufio.NewScanner(input)
  for sc.Scan() {
    line := sc.Text()

    if line == "" {
      continue
    }
    sack_len := len(line)
    first_sack := line[:sack_len/2]
    second_sack := line[sack_len/2:]

    for _, item := range first_sack {
        if strings.Contains(second_sack, string(item)) {
        fmt.Printf("%d %s Saco 1 e 2 contem %s\n", l, line, string(item))
      }
    }
    l++
  }
}
