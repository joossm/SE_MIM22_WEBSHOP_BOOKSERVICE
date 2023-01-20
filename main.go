package main

import (
	"SE_MIM22_WEBSHOP_BOOKSERVICE/handler"
	"log"
	"net/http"
	"time"
)

func main() {
	var serveMux = http.NewServeMux()
	serveMux.HandleFunc("/getAllBooks", handler.GetAllBooks)
	serveMux.HandleFunc("/getBookById", handler.GetBookByID)
	log.Printf("\n\n\tBOOKSERVICE\n\nAbout to listen on Port: 8440." +
		"\n\nSUPPORTED REQUESTS:\nGET:" +
		"\nGet All Books: http://127.0.0.1:8440/getAllBooks" +
		"\nGet Book By ID: http://127.0.0.1:8440/getBookById?id=1 requiers a url parameter id")
	server := &http.Server{
		Addr:              ":8440",
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       3 * time.Second,
		Handler:           serveMux,
	}
	log.Fatal(server.ListenAndServe())
}
