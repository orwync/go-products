// Package classification Category API.
//
// Documentation of our Category API.
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta
package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/orwync/go-products/helper"
)

func HandleRoutes() *mux.Router {
	r := mux.NewRouter()

	//Category api
	r.HandleFunc("/category", GetAllCategories).Methods(http.MethodGet)
	r.HandleFunc("/category/{id:[0-9]+}", GetCategory).Methods(http.MethodGet)
	r.HandleFunc("/category", CreateCategory).Methods(http.MethodPost)
	r.HandleFunc("/category/{id:[0-9]+}", UpdateCategory).Methods(http.MethodPut)
	r.HandleFunc("/category/{id:[0-9]+}", DeleteCategory).Methods(http.MethodDelete)

	//Product api
	r.HandleFunc("/product", GetAllProducts).Methods(http.MethodGet)
	r.HandleFunc("/product/{id:[0-9]+}", GetProduct).Methods(http.MethodGet)
	r.HandleFunc("/product", CreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/product/{id:[0-9]+}", UpdateProduct).Methods(http.MethodPut)
	r.HandleFunc("/product/{id:[0-9]+}", DeleteProduct).Methods(http.MethodDelete)

	//Variant api
	r.HandleFunc("/variant", GetAllVariants).Methods(http.MethodGet)
	r.HandleFunc("/variant/{id:[0-9]+}", GetVariant).Methods(http.MethodGet)
	r.HandleFunc("/variant", CreateVariant).Methods(http.MethodPost)
	r.HandleFunc("/variant/{id:[0-9]+}", UpdateVariant).Methods(http.MethodPut)
	r.HandleFunc("/variant/{id:[0-9]+}", DeleteVariant).Methods(http.MethodDelete)

	//Not Found
	r.NotFoundHandler = http.HandlerFunc(notFound)

	return r

}

func notFound(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusNotFound)
	errorMessage := helper.ErrorMessage("Page not found")
	json.NewEncoder(rw).Encode(errorMessage)
}
