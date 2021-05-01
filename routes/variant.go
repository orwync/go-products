package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/orwync/go-products/data"
	"github.com/orwync/go-products/helper"
)

func GetAllVariants(rw http.ResponseWriter, r *http.Request) {
	var variant []data.Variant
	data.DBConn.Find(&variant)
	err := json.NewEncoder(rw).Encode(variant)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func GetVariant(rw http.ResponseWriter, r *http.Request) {
	var variant data.Variant
	id := mux.Vars(r)["id"]
	err := data.DBConn.First(&variant, id).Error

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}
	rw.WriteHeader(200)
	err = json.NewEncoder(rw).Encode(variant)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func CreateVariant(rw http.ResponseWriter, r *http.Request) {
	var variant data.Variant
	err := json.NewDecoder(r.Body).Decode(&variant)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	result := data.DBConn.Create(&variant)

	fmt.Println(result.Error)
	json.NewEncoder(rw).Encode(variant)
}

func UpdateVariant(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Update Variant")
}

func DeleteVariant(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := data.DBConn.Delete(&data.Variant{}, id).Error

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	rw.WriteHeader(200)
	message := helper.SuccessMessage("variant deleted")
	json.NewEncoder(rw).Encode(message)
}
