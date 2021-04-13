package config

import (   
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
  
  "fmt"
)

var (
  db *gorm.DB
  err error
)

// DB接続
func Connect() *gorm.DB {

  dbms := "mysql"
  database := "database"
  user := "user"
  pass := "password"
  protcol := "tcp(db)"
  con := user + ":" + pass + "@" + protcol + "/" + database + "?parseTime=true"

  db, err := gorm.Open(dbms, con)
  if err != nil {
    fmt.Sprintf("DB接続に失敗しました。")
  }
  
  return db
}

// DB終了
func Close() {
  if err := db.Close(); err != nil {
    panic(err)
  }
}