package config

import (
	// "os"
   
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
  // "github.com/joho/godotenv"
  
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

  // // 実行環境取得
  // env := os.Getenv("ENV")
    
  // if "production" == env {
  //   env = "production"
  // } else {
  //   env = "development"
  // }

  // // 環境変数取得
  // godotenv.Load(".env." + env)
  // godotenv.Load()
    
  // // DB接続
  // db, err = gorm.Open("mysql", os.Getenv("CONNECT"))
    
  // if err != nil {
  //   panic(err)
  // }
  
  return db
}

// DB終了
func Close() {
  if err := db.Close(); err != nil {
    panic(err)
  }
}