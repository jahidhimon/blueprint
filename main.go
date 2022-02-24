package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
// https://youtu.be/SonwZ6MF5BE?t=2172
func main() {
	router := mux.NewRouter() // Router variable
	// Handlers
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books{id}", deleteBook).Methods("DELETE")
	// Listen to http server
	log.Fatal(http.ListenAndServe(":8000", router))
}

func pain() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html><head>Hello world</head><body>Let's Chat</body></html>`))
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("Server started")
}
