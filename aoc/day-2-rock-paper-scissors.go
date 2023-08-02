package main

import (
  "aoc/utils"
  "bufio"
  "fmt"
)

func main() {
  input := utils.ReadInputs("inputs/day-2.txt");
  defer input.Close();

  scores := map[string]int{
    "B X": 1,
    "C Y": 2,
    "A Z": 3,
    "A X": 4,
    "B Y": 5,
    "C Z": 6,
    "C X": 7,
    "A Y": 8,
    "B Z": 9,
  }
  b_scores := map[string]int {
    "A X": 3,
    "A Y": 4,
    "A Z": 8,
    "B X": 1,
    "B Y": 5,
    "B Z": 9,
    "C X": 2,
    "C Y": 6,
    "C Z": 7,
  }

  total := 0 
  b_total := 0

  sc := bufio.NewScanner(input);
  for sc.Scan() {
    line := sc.Text();

    if score, ok := scores[line]; ok {
      total += score;
    }

    if score, ok := b_scores[line]; ok {
      b_total += score;
    }
  }

  fmt.Printf("the final score is %d\npart two final score is %d", total, b_total)
}
