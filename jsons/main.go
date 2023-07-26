package main

import (
	"bytes"
  "encoding/json"
  "fmt"
	"log"
)

type Dog struct {
  Name string `json:"name"`
  Breed string `json:"breed"`
  Age uint `json:"age"`
}

func printJson(input_struct any) {
  struct_json, err := json.Marshal(input_struct);

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(bytes.NewBuffer(struct_json))
}

func main() {
  // Transforma o Struct em Json: json.Marshal() 
  // Transforma o Json em Struct: json.Unmarshal() 

vim.keymap.silent("t", "<esc>", "<C-\\><C-N>", t_opts)
  my_dog := Dog{"Rex", "Dalmata", 3}

  my_new_dog := map[string]string{
    "nome": "Cebola",
    "raca": "Golden Retriever",
  }

  printJson(my_dog)
  printJson(my_new_dog)
} 

