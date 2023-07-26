package main

import (
	"database/sql"
	"fmt"
  "log"

  _ "github.com/go-sql-driver/mysql"
)

func main() {
  connection_string := "user:password@/gui?charset=utf8&parseTime=True&loc=Local"

  db, err := sql.Open("mysql", connection_string)
  
  if err != nil {
    log.Fatal(err)
  }

  defer db.Close()

  fmt.Println(db)

  if err = db.Ping(); err != nil {
    log.Fatal(err)
  }
 
  fmt.Println("Conexão está aberta!")

  rows, err := db.Query("SELECT * FROM usuarios")

  if err != nil {
    log.Fatal(err)
  }

  defer rows.Close()

  fmt.Println(rows)
}

