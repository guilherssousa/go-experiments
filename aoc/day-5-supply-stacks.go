package main

import (
  "aoc/utils"
  "bufio"
  "fmt"
  "strconv"
  "strings"
)

type Stack struct {
  data []string
}

func (s *Stack) Push(item... string) {
  s.data = append(s.data, item...)
}

func (s Stack) Peek() string {
  return s.data[len(s.data)-1]
}

func (s *Stack) Pop() string {
  item := s.data[len(s.data)-1]
  s.data = s.data[:len(s.data)-1]
  return item
}

type Instruction struct {
  from int64
  to int64
  amount int64
}

func parseInstruction(str string) Instruction {
  data := strings.Split(str, " ")

  from, err := strconv.ParseInt(data[3], 10, 64)
  utils.Check(err)

  to, err := strconv.ParseInt(data[5], 10, 64)
  utils.Check(err)

  amount, err := strconv.ParseInt(data[1], 10, 64)
  utils.Check(err)

  return Instruction{from-1,to-1,amount}
}

func main() {
  input := utils.ReadInputs("inputs/day-5.txt");
  defer input.Close();

  sc := bufio.NewScanner(input);

  supply := []Stack{
    Stack{
      []string{"D", "T", "R", "B", "J", "L", "W", "G"},
    },
    Stack{
      []string{"S", "W", "C"},
    },
    Stack{
      []string{"R", "Z", "T", "M"},
    },
    Stack{
      []string{"D", "T", "C", "H", "S", "P", "V"},
    },
    Stack{
      []string{"G", "P", "T", "L", "D", "Z"},
    },
    Stack{
      []string{"F", "B", "R", "Z", "J", "Q", "C", "D"},
    },
    Stack{
      []string{"S", "B", "D", "J", "M", "F", "T", "R"},
    },
    Stack{
      []string{"L", "H", "R", "B", "T", "V", "M"},
    },
    Stack{
      []string{"Q", "P", "D", "S", "V"},
    },
  }

  supply_b := make([]Stack, len(supply))
  copy(supply_b, supply)

  for sc.Scan() {
    line := sc.Text()
    instruction := parseInstruction(line)

    for i := 0; i < int(instruction.amount); i++ {
      item := supply[instruction.from].Pop()
      supply[instruction.to].Push(item)
    }
  }

  final := ""
  for _, l := range supply {
    final += l.Peek()
  }
  fmt.Println("part 1 answer", final)

  input = utils.ReadInputs("inputs/day-5.txt")
  defer input.Close()

  bc := bufio.NewScanner(input)
  for bc.Scan() {
    line := bc.Text()
    instruction := parseInstruction(line)

    fmt.Println(line, supply_b[instruction.from], supply_b[instruction.to])
    to_move := []string{}
    for i := 0; i < int(instruction.amount); i++ {
      item := supply_b[instruction.from].Pop()
      to_move = append([]string{item}, to_move...)
    }
    supply_b[instruction.to].Push(to_move...)
    fmt.Println("finished", supply_b[instruction.from], supply_b[instruction.to])
  }

  final = ""
  for _, l := range supply_b {
    final += l.Peek()
  }

  fmt.Println("part 2 answer", final)
}
