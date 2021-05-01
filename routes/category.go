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

// swagger:route GET /category category listCategory
// Returns a list category
// responses:
// 200: categoriesResponse
// 500: statusMessage
func GetAllCategories(rw http.ResponseWriter, r *http.Request) {
	categories, err := services.GetAllCategories()

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}
	err = json.NewEncoder(rw).Encode(categories)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}
}

// swagger:route GET /category/{id} category getCategory
// Returns a category by id
// responses:
// 200: Category
// 500: statusMessage
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

// swagger:route POST /category category createCategory
// Creates a new category
// parameters: createCategory
// responses:
// 200: Category
// 500: statusMessage
func CreateCategory(rw http.ResponseWriter, r *http.Request) {
	var category data.Category
	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	err = services.CreateCategory(&category)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	json.NewEncoder(rw).Encode(category)

}

// swagger:route PUT /category/{id} category editCategory
// Edit existing category
// responses:
// 200: Category
// 500: statusMessage
func UpdateCategory(rw http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var updateCategory data.Category

	err := json.NewDecoder(r.Body).Decode(&updateCategory)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	err = services.UpdateCategory(id, &updateCategory)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	rw.WriteHeader(200)
	message := helper.SuccessMessage("Category updated")
	json.NewEncoder(rw).Encode(message)
}

// swagger:route DELETE /category/{id} category deleteCategory
// Delete existing category
// responses:
// 200: statusMessage
// 500: statusMessage
func DeleteCategory(rw http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := services.DeleteCategory(id)

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
