package handler

import (
	"SE_MIM22_WEBSHOP_BOOKSERVICE/model"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // MySQL Driver
	"net/http"
)

func GetAllBooks(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		db := openDB()
		defer closeDB(db)
		result, err := db.Query("SELECT * FROM books")
		errorHandler(err)
		var books []model.Book
		if result != nil {
			for result.Next() {
				var book model.Book
				err = result.Scan(&book.Id, &book.Titel, &book.EAN, &book.Content, &book.Price)
				errorHandler(err)
				books = append(books, book)
			}
		}
		jsonBook, err := json.Marshal(books)
		errorHandler(err)
		_, responseErr := responseWriter.Write(jsonBook)
		errorHandler(responseErr)
		return
	default:
		js, err := json.Marshal("THIS IS A GET REQUEST")
		errorHandler(err)
		_, responseErr := responseWriter.Write(js)
		errorHandler(responseErr)
		return
	}
}

func GetBookByID(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		db := openDB()
		defer closeDB(db)
		result, err := db.Query("SELECT * FROM books WHERE Id = ?", request.URL.Query().Get("id"))
		errorHandler(err)
		var books []model.Book
		if result != nil {
			for result.Next() {
				var book model.Book
				err = result.Scan(&book.Id, &book.Titel, &book.EAN, &book.Content, &book.Price)
				errorHandler(err)
				books = append(books, book)
			}
		}
		jsonBook, err := json.Marshal(books[0])
		errorHandler(err)
		_, responseErr := responseWriter.Write(jsonBook)
		errorHandler(responseErr)
		return
	default:
		js, err := json.Marshal("THIS IS A GET REQUEST")
		errorHandler(err)
		_, responseErr := responseWriter.Write(js)
		errorHandler(responseErr)
		return
	}
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func closeDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		print(err)
	}
}

func openDB() *sql.DB {
	fmt.Println("Opening DB")
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/books")
	fmt.Println(db.Ping())
	fmt.Println(db.Stats())
	db.SetMaxIdleConns(0)
	errorHandler(err)
	return db
}
