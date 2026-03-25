package handlers

import (
	"encoding/json"
	"net/http"

	"bookstore/models"
)

var Categories []models.Category
var CategoryID = 1

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
 
	result := []models.Category{}
	result = append(result, Categories...)
	json.NewEncoder(w).Encode(result)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var category models.Category
	json.NewDecoder(r.Body).Decode(&category)

	if category.Name == "" {
		http.Error(w, "Invalid data", 400)
		return
	}

	category.ID = CategoryID
	CategoryID++
	Categories = append(Categories, category)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}