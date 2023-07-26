package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Connect opens conection with Database.
func Connect() (*sql.DB, error) {
  connection_str := "user:password@/gui?charset=utf8&parseTime=True&loc=Local"

  db, err := sql.Open("mysql", connection_str)
  if err != nil {
    return nil, err
  }

  if err = db.Ping(); err != nil {
    return nil, err
  }

  return db, nil
} 
