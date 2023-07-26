package main

import (
	"log"
	"net/http"
)

func main() {

  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hELLO FUTURE"))
  })

  log.Fatal(http.ListenAndServe(":3000", nil))
}
