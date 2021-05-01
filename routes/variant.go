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

// swagger:route GET /variant variant listVariant
// Returns a list variant
// responses:
// 200: categoriesResponse
// 500: statusMessage
func GetAllVariants(rw http.ResponseWriter, r *http.Request) {
	variants, err := services.GetAllVariants()

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	err = json.NewEncoder(rw).Encode(variants)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}
}

// swagger:route GET /variant/{id} variant getVariant
// Returns a variant by id
// responses:
// 200: Variant
// 500: statusMessage
func GetVariant(rw http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	variant, err := services.GetVariant(id)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}
	rw.WriteHeader(200)
	err = json.NewEncoder(rw).Encode(variant)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}
}

// swagger:route POST /variant variant createVariant
// Creates a new variant
// responses:
// 200: Variant
// 500: statusMessage
func CreateVariant(rw http.ResponseWriter, r *http.Request) {
	var variant data.Variant
	err := json.NewDecoder(r.Body).Decode(&variant)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	err = services.CreateVariant(&variant)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	json.NewEncoder(rw).Encode(variant)
}

// swagger:route PUT /variant/{id} variant editVariant
// Edit existing variant
// responses:
// 200: Variant
// 500: statusMessage
func UpdateVariant(rw http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var updateVariant data.Variant
	err := json.NewDecoder(r.Body).Decode(&updateVariant)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	err = services.UpdateVariant(id, &updateVariant)

	if err != nil {
		rw.WriteHeader(500)
		errorMessage := helper.ErrorMessage(err.Error())
		json.NewEncoder(rw).Encode(errorMessage)
		return
	}

	rw.WriteHeader(200)
	message := helper.SuccessMessage("variant updated")
	json.NewEncoder(rw).Encode(message)
}

// swagger:route DELETE /variant/{id} variant deleteVariant
// Delete existing variant
// responses:
// 200: successMessage
// 500: statusMessage
func DeleteVariant(rw http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := services.DeleteVariant(id)

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
