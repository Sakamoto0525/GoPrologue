package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/joho/godotenv"
  
  "os"
  "log"
)

var (
  db *gorm.DB
  err error
)

// DB接続
func Connect() *gorm.DB {

  err := godotenv.Load()
  if err != nil {
      log.Fatal("Error loading .env file")
  }

  dbms := "mysql"
  database := os.Getenv("DB")
  user := os.Getenv("DB_USER")
  pass := os.Getenv("DB_PASS")
  protcol := "tcp(db)"
  con := user + ":" + pass + "@" + protcol + "/" + database + "?parseTime=true"

  db, err := gorm.Open(dbms, con)
  if err != nil {
    log.Fatal("DB接続に失敗しました。")
  }
  
  return db
}

// DB終了
func Close() {
  if err := db.Close(); err != nil {
    panic(err)
  }
}