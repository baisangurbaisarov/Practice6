package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"bookstore/models"
)

var Books []models.Book
var BookID = 1

func GetBooks(w http.ResponseWriter, r *http.Request) {

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	categoryStr := r.URL.Query().Get("category")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	categoryID, _ := strconv.Atoi(categoryStr)

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 5
	}

	var filtered []models.Book
	for _, b := range Books {
		if categoryStr != "" && b.CategoryID != categoryID {
			continue
		}
		filtered = append(filtered, b)
	}

	start := (page - 1) * limit
	end := start + limit

	if start > len(filtered) {
		start = len(filtered)
	}
	if end > len(filtered) {
		end = len(filtered)
	}

	json.NewEncoder(w).Encode(filtered[start:end])
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, _ := strconv.Atoi(idStr)

	for _, b := range Books {
		if b.ID == id {
			json.NewEncoder(w).Encode(b)
			return
		}
	}
	http.Error(w, "Not found", 404)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	if book.Title == "" || book.Price <= 0 {
		http.Error(w, "Invalid data", 400)
		return
	}

	book.ID = BookID
	BookID++
	Books = append(Books, book)

	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, _ := strconv.Atoi(idStr)

	for i := range Books {
		if Books[i].ID == id {
			json.NewDecoder(r.Body).Decode(&Books[i])
			Books[i].ID = id
			json.NewEncoder(w).Encode(Books[i])
			return
		}
	}
	http.Error(w, "Not found", 404)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, _ := strconv.Atoi(idStr)

	for i := range Books {
		if Books[i].ID == id {
			Books = append(Books[:i], Books[i+1:]...)
			return
		}
	}
	http.Error(w, "Not found", 404)
}