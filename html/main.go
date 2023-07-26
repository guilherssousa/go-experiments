package main

import (
  "log"
  "net/http"
  "html/template"
)

type usuario struct {
  Nome string
  Email string 
}

var templates *template.Template

func main() {
  templates = template.Must(template.ParseGlob("*.html"))

  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    u:= usuario{
      "Guilherme",
      "gui@guisousa.com",
    }
    templates.ExecuteTemplate(w, "index.html",u) 
  })

  log.Fatal(http.ListenAndServe(":3000", nil))
}
