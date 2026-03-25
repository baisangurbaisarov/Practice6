package handlers

import (
	"encoding/json"
	"net/http"

	"bookstore/models"
)

var Categories []models.Category
var CategoryID = 1

func GetCategories(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	json.NewDecoder(r.Body).Decode(&category)

	if category.Name == "" {
		http.Error(w, "Invalid data", 400)
		return
	}

	category.ID = CategoryID
	CategoryID++
	Categories = append(Categories, category)

	json.NewEncoder(w).Encode(category)
}