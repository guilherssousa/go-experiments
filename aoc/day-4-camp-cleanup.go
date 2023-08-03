package main

import (
  "aoc/utils"
  "bufio"
  "fmt"
  "strconv"
  "strings"
)

const sections string = "123456789"

func contains(a, b, c, d int) bool {
  result := (a <= c) && (b >= d) 
  return result
}

func main() {
  input := utils.ReadInputs("inputs/day-4.txt")
  defer input.Close()

  covered_ranged_pairs := 0

  sc := bufio.NewScanner(input)
  for sc.Scan() {
    line := sc.Text()

    pairs := strings.Split(line, ",")
    pair_a := []int64{}
    pair_b := []int64{}

    for i, pair := range pairs {
      str_pair := strings.Split(pair, "-")
      for _, v := range str_pair { 
        iv, err := strconv.ParseInt(v, 10, 64)
        utils.Check(err)

        if i == 0 {
          pair_a = append(pair_a, iv)
        } else {
          pair_b = append(pair_b, iv)
        }
      }
    }

    a_contains_b := contains(int(pair_a[0]), int(pair_a[1]), int(pair_b[0]), int(pair_b[1]))
    if a_contains_b {
      covered_ranged_pairs++
      continue
    }

    b_contains_a := contains(int(pair_b[0]), int(pair_b[1]), int(pair_a[0]), int(pair_a[1]))
    if b_contains_a {
      covered_ranged_pairs++
    }
  }

  fmt.Println(covered_ranged_pairs)
}
