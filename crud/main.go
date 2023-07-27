package main

import (
  "crud/server"
  "fmt"
  "html/template"
  "log"
  "net/http"

  "github.com/joho/godotenv"
  "github.com/gorilla/mux"
)

func main() {
  // CRUD é Create-Read-Update-Delete.
  // Não sei porque estou escrevendo isso denovo pela 40323a vez?
  err := godotenv.Load() 
  if err != nil {
    log.Fatal("error loading .env file", err)
  }

  router := mux.NewRouter()
  engine := template.Must(template.ParseGlob("views/*.html"))

  // API Routes
  router.HandleFunc("/usuarios", server.CreateUser).Methods(http.MethodPost)
  router.HandleFunc("/usuarios", server.ListUsers).Methods(http.MethodGet)
  router.HandleFunc("/usuarios/{id}", server.GetUser).Methods(http.MethodGet)
  router.HandleFunc("/usuarios/{id}", server.UpdateUser).Methods(http.MethodPut)
  router.HandleFunc("/usuarios/{id}", server.DeleteUser).Methods(http.MethodDelete)

  // View routes
  router.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    engine.ExecuteTemplate(w, "index.html", nil)
  }).Methods(http.MethodGet)

  // Static file serving
  router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

  fmt.Println("Escutando em http://localhost:5000")
  log.Fatal(http.ListenAndServe(":5000", router))
}
