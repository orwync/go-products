package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/orwync/go-products/data"
	"github.com/orwync/go-products/helper"
)

func GetAllProducts(rw http.ResponseWriter, r *http.Request) {
	var products []data.Product
	data.DBConn.Find(&products)
	err := json.NewEncoder(rw).Encode(products)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func GetProduct(rw http.ResponseWriter, r *http.Request) {
	var product data.Product
	id := mux.Vars(r)["id"]
	err := data.DBConn.First(&product, id).Error

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

func UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Update Product")
}

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
	message := helper.SuccessMessage("category deleted")
	json.NewEncoder(rw).Encode(message)
}
