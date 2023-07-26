package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
  // You can use `json:"-"` if you want to
  // ignore a JSON key during json.Unmarshal
  Nome string `json:"nome"`
  Raca string `json:"raca"`
  Idade uint `json:"idade"`
}

func main() {
  stringDog := `{"nome": "rex", "raca": "Dálmata", "idade": 3}`
  var c cachorro

  if erro := json.Unmarshal([]byte(stringDog), &c); erro != nil {
    log.Fatal(erro)
  }

  // You can do the same with map[string]string

  fmt.Println(c)
}
