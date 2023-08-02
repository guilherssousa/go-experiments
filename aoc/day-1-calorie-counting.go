package main

import (
  "bufio"
  "fmt"
  "sort"
  "strconv"
  "aoc/utils"
)

func main() {
  input := utils.ReadInputs("inputs/day-1.txt");
  defer input.Close();

  elfs := []int{}
  var accumulated int 
  var currentElf int

  sc := bufio.NewScanner(input);
  for sc.Scan() {
    line := sc.Text();

    if line == "" {
      elfs = append(elfs, accumulated);
      currentElf++;
      accumulated = 0;
      continue;
    }

    v, err := strconv.ParseInt(line, 10, 64);
    utils.Check(err)

    accumulated += int(v)
  }
  sort.Ints(elfs)
  
  top3 := 0 

  for _, v := range elfs[len(elfs)-3:] {
    fmt.Println(v)
    top3 += v
  }

  fmt.Println("the top3 sum is", top3)
}
