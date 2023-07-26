package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// HTTP é um protocolo de comunicação.
	// Ele é a base da comunicação Web.
	// Ele funciona no relacionamento Cliente-Servidor

  health := true
  
  http.HandleFunc("/ruok", func (w http.ResponseWriter, r *http.Request) {
    if health {
      w.Write([]byte("imok"))
      health = false
    } else {
      w.Write([]byte("imnotok"))
    }
  })

  http.HandleFunc("/home", func (w http.ResponseWriter, r *http.Request) {
    fmt.Println("Recebendo um Request")

    w.Write([]byte("Hello world")) 
  })

  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Hello world!</h1>"))
  })

  log.Fatal(http.ListenAndServe(":3000", nil))
}
