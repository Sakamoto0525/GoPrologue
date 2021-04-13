package main

import (
    "github.com/gorilla/mux"
	"log"
	"encoding/json"
    "net/http"
    "work/config"
    "io/ioutil"
    "github.com/jinzhu/gorm"
)

type Book struct {
    gorm.Model
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Author *Author `json:"author"`
}

type Author struct {
    gorm.Model
    FirstName string `json:"firstname"`
    LastName  string `json:"lastname"`
}

var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    db := config.Connect()
    books := db.Find(&books)

    json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
   
    params := mux.Vars(r)
    db := config.Connect()
    book := db.First(&books, params["id"])

    json.NewEncoder(w).Encode(book)
}

// Create a Book
func createBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    reqBody, _ := ioutil.ReadAll(r.Body)
    var book Book
    if err := json.Unmarshal(reqBody, &book); err != nil {
        log.Fatal(err)
    }

    db := config.Connect()
    db.NewRecord(book)
    db.Create(&book)

    json.NewEncoder(w).Encode(book)
}

// Update a Book
func updateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    reqBody, _ := ioutil.ReadAll(r.Body)

    var book Book
    if err := json.Unmarshal(reqBody, &book); err != nil {
        log.Fatal(err)
    }

    db := config.Connect()
    db.Model(&book).Where("id = ?", params["id"]).Update(&book)

    json.NewEncoder(w).Encode(books)
}

// Delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    db := config.Connect()
    db.Delete(&books, params["id"])
}

func main() {
    // ルーターのイニシャライズ
	r := mux.NewRouter()

    db := config.Connect()
    db.AutoMigrate(&Book{})
    db.AutoMigrate(&Author{})

    // ルート(エンドポイント)
    r.HandleFunc("/api/books", getBooks).Methods("GET")
    r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
    r.HandleFunc("/api/books", createBook).Methods("POST")
    r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
    r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8080", r))
}
