package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "sort"
  "strconv"
)

func check(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func main() {
  input, err := os.OpenFile("inputs/day-1.txt", os.O_RDONLY, os.ModePerm);
  check(err)
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
    check(err)

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
