package routes

import (
	"fmt"
	"net/http"
)

func GetAllProducts(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Get categories endpoint")
}

func GetProduct(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Get Product")
}

func CreateProduct(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Create Product")
}

func UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Update Product")
}

func DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Delete Product")
}
