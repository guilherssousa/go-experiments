package server

import (
  "crud/database"
  "encoding/json"
  "fmt"
  "io/ioutil"

  "net/http"
)

type user struct {
  ID uint32 `json:"id"`
  Nome string `json:"nome"`
  Email string `json:"email"`
}

// CreateUser inserts a new User on DATABASE
func CreateUser(w http.ResponseWriter, r *http.Request) {
  requestBody, err := ioutil.ReadAll(r.Body)  

  if err != nil {
    w.Write([]byte("Falha ao ler o corpo da requisição."))
    return
  }

  var user user 

  if err = json.Unmarshal(requestBody, &user); err != nil {
    w.Write([]byte("Erro ao converter o corpo da requisição"))    
    return
  }

  db, err := database.Connect()
  if err != nil {
    w.Write([]byte("Erro ao se conectar com o banco de dados."))
    return
  }
  defer db.Close()  

  // INSERT INTO usuarios (nome, email) values ("nome", "email")

  // PREPARE STATEMENT
  statement, err :=  db.Prepare("INSERT INTO usuarios (nome, email) values (?, ?)")
  if err != nil {
    w.Write([]byte("Erro ao criar o statement!"))
    return
  }
   defer statement.Close()

   insert, err := statement.Exec(user.Nome, user.Email)
  
  if err != nil {
    w.Write([]byte("Erro ao executar o statement!"))
    return
  }

  insertId, err := insert.LastInsertId()
  if err != nil {
    w.Write([]byte("Erro ao obter o ID inserido"))
    return
  }
  
  w.WriteHeader(http.StatusCreated)
  w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", insertId)))
}
