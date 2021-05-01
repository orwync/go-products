package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/orwync/go-products/data"
	"github.com/orwync/go-products/helper"
	"github.com/orwync/go-products/services"
)

func GetAllCategories(rw http.ResponseWriter, r *http.Request) {
	categories, err := services.GetAllCategories()

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(rw).Encode(categories)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func GetCategory(rw http.ResponseWriter, r *http.Request) {
	var category data.Category
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	category, err := services.GetCategory(id)

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

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	err = data.DBConn.Create(&category).Error

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	json.NewEncoder(rw).Encode(category)

}

func UpdateCategory(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var category data.Category
	err := data.DBConn.Find(&category, id).Error

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	var updateCategory data.Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	data.DBConn.Model(&category).Updates(updateCategory)

	rw.WriteHeader(200)
	message := helper.SuccessMessage("Category updated")
	json.NewEncoder(rw).Encode(message)
}

func DeleteCategory(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := data.DBConn.Delete(&data.Category{}, id).Error

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	rw.WriteHeader(200)
	message := helper.SuccessMessage("category deleted")
	json.NewEncoder(rw).Encode(message)
}
