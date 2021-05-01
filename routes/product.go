package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/orwync/go-products/data"
	"github.com/orwync/go-products/helper"
)

// swagger:route GET /product product listProduct
// Returns a list product
// responses:
// 200: categoriesResponse
// 500: statusMessage
func GetAllProducts(rw http.ResponseWriter, r *http.Request) {
	var products []data.Product
	data.DBConn.Preload("Variant").Find(&products)
	err := json.NewEncoder(rw).Encode(products)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

// swagger:route GET /product/{id} product getProduct
// Returns a product by id
// responses:
// 200: Product
// 500: statusMessage
func GetProduct(rw http.ResponseWriter, r *http.Request) {
	var product data.Product
	id := mux.Vars(r)["id"]
	err := data.DBConn.Preload("Variant").First(&product, id).Error

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}
	rw.WriteHeader(200)
	err = json.NewEncoder(rw).Encode(product)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

// swagger:route POST /product product createProduct
// Creates a new product
// responses:
// 200: Product
// 500: statusMessage
func CreateProduct(rw http.ResponseWriter, r *http.Request) {
	var product data.Product
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	result := data.DBConn.Create(&product)

	fmt.Println(result.Error)
	json.NewEncoder(rw).Encode(product)
}

// swagger:route PUT /product/{id} product editProduct
// Edit existing product
// responses:
// 200: Product
// 500: statusMessage
func UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var product data.Product
	err := data.DBConn.Find(&product, id).Error

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	var updateProduct data.Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	data.DBConn.Model(&product).Updates(updateProduct)

	rw.WriteHeader(200)
	message := helper.SuccessMessage("product updated")
	json.NewEncoder(rw).Encode(message)
}

// swagger:route DELETE /product/{id} product deleteProduct
// Delete existing product
// responses:
// 200: successMessage
// 500: errorMessage
func DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := data.DBConn.Delete(&data.Product{}, id).Error

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	rw.WriteHeader(200)
	message := helper.SuccessMessage("product deleted")
	json.NewEncoder(rw).Encode(message)
}
