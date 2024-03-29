package server

import (
  "crud/database"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "strconv"
  "time"

  "github.com/gorilla/mux"
)

type message struct {
  ID uint32 `json:"id,omitempty"`
  AuthorId int `db:"author_id" json:"author_id,omitempty"`
  PublishedAt *time.Time `db:"published_at" json:"published_at"`
  Title string `json:"title"`
  Content string `db:"content" json:"content,omitempty"`
  user
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
  requestBody, err := ioutil.ReadAll(r.Body)  

  if err != nil {
    log.Print(err)
    w.Write([]byte("Falha ao ler o corpo da requisição."))
    return
  }

  var message message 

  if err = json.Unmarshal(requestBody, &message); err != nil {
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

  // PREPARE STATEMENT
  statement, err :=  db.Prepare("INSERT INTO messages (title, content, author_id) values ($1, $2, 6)")
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao criar o statement!"))
    return
  }

  _, err = statement.Exec(message.Title, message.Content)
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao executar o statement!"))
    return
  }

  w.WriteHeader(http.StatusCreated)
  w.Write([]byte(fmt.Sprintf("Mensagem inserida com sucesso!")))
}

func ListMessages(w http.ResponseWriter, r *http.Request) {
  db, err := database.Connect()
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao conectar com o banco de dados!"))
    return
  }
  defer db.Close()

  sql_query := "SELECT messages.id, messages.title, messages.published_at, usuarios.nome as name, usuarios.email FROM messages LEFT JOIN usuarios ON author_id = usuarios.id ORDER BY published_at DESC"
  rows, err := db.Query(sql_query)
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao buscar mensagens"))
    return
  }
  defer rows.Close()

  var messages []message

  for rows.Next() {
    var message message

    if err := rows.Scan(
      &message.ID,
      &message.Title,
      &message.PublishedAt, 
      &message.user.Nome,
      &message.user.Email,
    ); err != nil {
      log.Println(err)
      w.Write([]byte("Erro ao escanear a mensagem."))
      return
    }

    messages = append(messages, message)
  }

  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(messages); err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao converter mensagens para JSON."))
    return
  }
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
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

  sql_statement := "SELECT messages.title, messages.content, messages.published_at, usuarios.nome, usuarios.email FROM messages LEFT JOIN usuarios ON usuarios.id = messages.author_id WHERE messages.id = $1"
  row, err := db.Query(sql_statement, ID)
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao buscar a mensagem")) 
    return 
  }

  var message message
  if row.Next() {
    if err := row.Scan(
        &message.Title,
        &message.Content,
        &message.PublishedAt,
        &message.user.Nome,
        &message.user.Email,
      ); err != nil {
      log.Println(err)
      w.Write([]byte("Erro ao mapear dados para a mensagem"))
      return
    }
  } 

  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(message); err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao converter a mensagem para JSON!"))
    return
  } 
}


