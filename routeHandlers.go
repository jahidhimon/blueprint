package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

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
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100001))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
	fmt.Println("A new book created")
}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}
