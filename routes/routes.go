package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello there")
	})

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

	return r

}
