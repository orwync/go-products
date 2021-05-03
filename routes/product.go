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

// swagger:route GET /product product listProduct
// Returns a list product
// responses:
// 200: productsResponse
// 500: statusMessage
func GetAllProducts(rw http.ResponseWriter, r *http.Request) {
	products, err := services.GetAllProducts()

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}
	err = json.NewEncoder(rw).Encode(products)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}
}

// swagger:route GET /product/{id} product getProduct
// Returns a product by id
// responses:
// 200: productResponse
// 500: statusMessage
func GetProduct(rw http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	product, err := services.GetProduct(id)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}
	rw.WriteHeader(200)
	err = json.NewEncoder(rw).Encode(product)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
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

	err = services.CreateProduct(&product)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	json.NewEncoder(rw).Encode(product)
}

// swagger:route PUT /product/{id} product editProduct
// Edit existing product
// responses:
// 200: Product
// 500: statusMessage
func UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var updateProduct data.Product
	err := json.NewDecoder(r.Body).Decode(&updateProduct)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	err = services.UpdateProduct(id, &updateProduct)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

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
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := services.DeleteProduct(id)

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
