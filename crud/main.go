package main

import (
  "crud/server"
  "fmt"
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

func main() {
  // CRUD é Create-Read-Update-Delete.
  // Não sei porque estou escrevendo isso denovo pela 40323a vez?

  router := mux.NewRouter()
  router.HandleFunc("/usuarios", server.CreateUser).Methods(http.MethodPost)
  router.HandleFunc("/usuarios", server.ListUsers).Methods(http.MethodGet)
  router.HandleFunc("/usuario/{id}", server.GetUser).Methods(http.MethodGet)

  fmt.Println("Escutando em http://localhost:5000")
  log.Fatal(http.ListenAndServe(":5000", router))
}
