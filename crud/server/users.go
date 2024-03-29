package server

import (
  "crud/database"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "strconv"
  "log"

  "net/http"

  "github.com/gorilla/mux"
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
    log.Print(err)
    w.Write([]byte("Falha ao ler o corpo da requisição."))
    return
  }

  var user user 

  if err = json.Unmarshal(requestBody, &user); err != nil {
    log.Print(err)
    w.Write([]byte("Erro ao converter o corpo da requisição"))    
    return
  }

  db, err := database.Connect()
  if err != nil {
    log.Print(err)
    w.Write([]byte("Erro ao se conectar com o banco de dados."))
    return
  }
  defer db.Close()  

  // INSERT INTO usuarios (nome, email) values ("nome", "email")

  // PREPARE STATEMENT
  statement, err :=  db.Prepare("INSERT INTO usuarios (nome, email) values ($1, $2)")
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao criar o statement!"))
    return
  }

  _, err = statement.Exec(user.Nome, user.Email)
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao executar o statement!"))
    return
  }

  w.WriteHeader(http.StatusCreated)
  w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso!")))
}

// SearchUsers list all Users
func ListUsers(w http.ResponseWriter, r *http.Request) {
  db, err := database.Connect()
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao conectar com o banco de dados!"))
    return
  }
  defer db.Close()

  rows, err := db.Query("SELECT * FROM usuarios")
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao buscar os usuários"))
    return
  }
  defer rows.Close()

  var users []user

  for rows.Next() {
    var user user

    if err := rows.Scan(&user.ID, &user.Nome, &user.Email); err != nil {
      log.Println(err)
      w.Write([]byte("Erro ao escanear o usuário."))
      return
    }

    users = append(users, user)
  }

  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(users); err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao converter os usuários para JSON."))
    return
  }
}

func GetUser(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)

  ID, err := strconv.ParseUint(params["id"], 10, 32)
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao converter os dados da requisição."))
    return
  }

  db, err := database.Connect()
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao conectar com o banco de dados!"))
    return
  }
  defer db.Close()

  row, err := db.Query("SELECT * FROM usuarios WHERE id = $1", ID)
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao buscar o usuário")) 
    return 
  }

  var user user
  if row.Next() {
    if err := row.Scan(&user.ID, &user.Nome, &user.Email); err != nil {
      log.Println(err)
      w.Write([]byte("Erro ao mapear dados para o usuário"))
      return
    }
  } 

  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(user); err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao converter o usuário para JSON!"))
    return
  } 
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)

  ID, err := strconv.ParseUint(params["id"], 10, 32)
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao converter os dados da requisição."))
    return
  }

  requestBody, err := ioutil.ReadAll(r.Body)
  if err != nil {
    log.Println(err)
    w.Write([]byte("Falha ao ler o corpo da requisição."))
    return
  }

  var user user
  if err = json.Unmarshal(requestBody, &user); err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao converter o corpo da requisição"))    
    return
  }

  db, err := database.Connect();
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao abrir conexão com o banco de dados!"))
    return
  }
  defer db.Close()

  statement, err := db.Prepare("UPDATE usuarios SET nome = $1, email = $2 WHERE id = $3")
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao criar o statement!"))
    return
  }
  defer statement.Close()

  if _, err := statement.Exec(user.Nome, user.Email, ID); err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao atualizar o usuário"))
    return
  }

  w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)

  ID, err := strconv.ParseUint(params["id"], 10, 32)
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao converter os dados da requisição."))
    return
  }

  db, err := database.Connect()
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao conectar com o banco de dados."))
    return
  }
  defer db.Close()

  statement, err := db.Prepare("DELETE FROM usuarios WHERE id = $1")
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao criar o statement!"))
    return
  }
  defer statement.Close()

  if _, err := statement.Exec(ID); err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao deletar o usuário."))
    return
  }

  w.WriteHeader(http.StatusNoContent)
}
