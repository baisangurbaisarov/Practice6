package handlers

import (
	"encoding/json"
	"net/http"

	"bookstore/models"
)

var Authors []models.Author
var AuthorID = 1

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Authors)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	json.NewDecoder(r.Body).Decode(&author)

	if author.Name == "" {
		http.Error(w, "Invalid data", 400)
		return
	}

	author.ID = AuthorID
	AuthorID++
	Authors = append(Authors, author)

	json.NewEncoder(w).Encode(author)
	
}