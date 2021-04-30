package routes

import (
	"fmt"
	"net/http"
)

func GetAllVariants(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Get categories endpoint")
}

func GetVariant(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Get Variant")
}

func CreateVariant(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Create Variant")
}

func UpdateVariant(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Update Variant")
}

func DeleteVariant(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Delete Variant")
}
