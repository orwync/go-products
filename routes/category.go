package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/orwync/go-products/data"
)

func GetAllCategories(rw http.ResponseWriter, r *http.Request) {
	var categories []data.Category
	data.DBConn.Find(&categories)
	rw.Header().Set("content-type", "application/json")
	err := json.NewEncoder(rw).Encode(categories)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func GetCategory(rw http.ResponseWriter, r *http.Request) {
	var category data.Category
	id := mux.Vars(r)["id"]
	result := data.DBConn.First(&category, id)

	if category.Name == "" {
		rw.WriteHeader(500)
		errorMessage := errorMessage(result.Error.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	err := json.NewEncoder(rw).Encode(category)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func CreateCategory(rw http.ResponseWriter, r *http.Request) {
	var category data.Category
	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := errorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	result := data.DBConn.Create(&category)

	fmt.Println(result.Error)
	fmt.Fprintf(rw, "Create Category")

}

func UpdateCategory(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Update Category")
}

func DeleteCategory(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Delete category")
}

func errorMessage(message string) data.Error {
	return data.Error{
		Status:  false,
		Message: message,
	}
}
