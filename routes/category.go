package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/orwync/go-products/data"
	"github.com/orwync/go-products/helper"
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
	err := data.DBConn.First(&category, id).Error
	rw.Header().Set("content-type", "application/json")

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}
	rw.WriteHeader(200)
	err = json.NewEncoder(rw).Encode(category)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func CreateCategory(rw http.ResponseWriter, r *http.Request) {
	var category data.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	rw.Header().Set("content-type", "application/json")

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	result := data.DBConn.Create(&category)

	fmt.Println(result.Error)
	json.NewEncoder(rw).Encode(category)

}

func UpdateCategory(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Update Category")
}

func DeleteCategory(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	data.DBConn.Delete(&data.Category{}, id)
	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(200)
	message := helper.SuccessMessage("category deleted")
	json.NewEncoder(rw).Encode(message)
}
