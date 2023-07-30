package main

import (
	"crud/server"
	"crud/util"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

  router.HandleFunc("/messages", server.CreateMessage).Methods(http.MethodPost)
  router.HandleFunc("/messages", server.ListMessages).Methods(http.MethodGet)
  router.HandleFunc("/messages/{id}", server.GetMessage).Methods(http.MethodGet)

  // View routes
  router.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    engine.ExecuteTemplate(w, "index.html", nil)
  }).Methods(http.MethodGet)

  router.HandleFunc("/new-message", func (w http.ResponseWriter, r *http.Request) {
    engine.ExecuteTemplate(w, "new-message.html", nil)
  }).Methods(http.MethodGet)

  router.HandleFunc("/message", func (w http.ResponseWriter, r *http.Request) {
    engine.ExecuteTemplate(w, "message.html", nil)
  }).Methods(http.MethodGet)

  // Static file serving
  router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

  localIPs, err := util.LocalIPs()
  if err != nil {
    log.Println(err)
  }

  fmt.Println("Escutando em http://localhost:5000")
  fmt.Println("Escutando em Host nos IPs:")
  for _, ip := range localIPs {
    fmt.Printf("\thttp://%s:5000\n", ip.To4().String())
  }

  log.Fatal(http.ListenAndServe(":5000", router))
}
