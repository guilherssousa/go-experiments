package database

import (
  "database/sql"
  "log"

  _ "github.com/lib/pq"
)

// Connect opens conection with Database.
func Connect() (*sql.DB, error) {
  //connection_str := "postgresql://user:password@/gui?charset=utf8&parseTime=True&loc=Local&sslmode=disable"
  connection_str := "user=user password=password dbname=gui sslmode=disable TimeZone=UTC"

  db, err := sql.Open("postgres", connection_str)
  if err != nil {
    log.Println(err)
    return nil, err
  }

  if err = db.Ping(); err != nil {
    log.Println(err)
    return nil, err
  }

  return db, nil
} 
