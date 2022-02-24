package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var desiredBook Book
	notFound := true
	for _, book := range books {
		if book.ID == params["id"] {
			desiredBook = book
			notFound = false
			break
		}
	}
	if notFound {
		error := Error{Code: "404", Message: "The book you are searching is not found on the server"}
		json.NewEncoder(w).Encode(error)
	} else {
		json.NewEncoder(w).Encode(desiredBook)
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}
