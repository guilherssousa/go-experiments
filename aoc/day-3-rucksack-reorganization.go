package main

import (
  "aoc/utils"
  "bufio"
  "fmt"
  "strings"
)

func findCommonItem(iterable string, toIterate ...string) string {
  common := string(iterable[0])
  loop1:
  for _, item := range iterable {
    for _, list := range toIterate {
      if !strings.Contains(list, string(item)) {
        continue loop1 
      }
    }
    common = string(item)
  }

  return common
}

func main() {
  input := utils.ReadInputs("inputs/day-3.txt")
  defer input.Close()

  priorities := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
  priority_sum := 0

  sc := bufio.NewScanner(input)
  sack_len := []string{}

  for sc.Scan() {
    line := sc.Text()

    if line == "" {
      continue
    }

    sack_len = append(sack_len, line)

    if len(sack_len) == 3 {
      common_item := findCommonItem(sack_len[0], sack_len[1:]...)
      priority_sum += strings.Index(priorities, common_item)
      sack_len = []string{} 
    }
  }

  fmt.Println("final:", priority_sum)
}
