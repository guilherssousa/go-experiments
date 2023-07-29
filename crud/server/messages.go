package server

import (
	"crud/database"
	"encoding/json"
	"log"
	"net/http"
  "time"
)

type message struct {
  ID uint32 `json:"id"`
  AuthorId int `db:"author_id" json:"author_id"`
  PublishedAt *time.Time `db:"published_at" json:"published_at"`
  Title string `json:"title"`
  Content string `db:"content" json:"content"`
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusCreated)
  return
}

func ListMessages(w http.ResponseWriter, r *http.Request) {
  db, err := database.Connect()
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao conectar com o banco de dados!"))
    return
  }
  defer db.Close()

  rows, err := db.Query("SELECT * FROM messages")
  if err != nil {
    log.Println(err)
    w.Write([]byte("Erro ao buscar mensagens"))
    return
  }
  defer rows.Close()

  var messages []message

  for rows.Next() {
    var message message

      if err := rows.Scan(&message.ID, &message.AuthorId, &message.Title, &message.Content, &message.PublishedAt); err != nil {
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
