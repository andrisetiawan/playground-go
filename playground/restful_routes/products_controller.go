package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// GET "/products/"
func GetProductsHandler(rw http.ResponseWriter, req *http.Request) {
	data := []Product{}
	data = append(data, Product{"1", "Foo"})
	data = append(data, Product{"2", "Bar"})

	products := Products{
		Data:  data,
		Total: len(data),
	}

	renderJSON(rw, http.StatusOK, products)
}

// POST "/products/"
func CreateProductsHandler(rw http.ResponseWriter, req *http.Request) {
	var product Product

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&product)
	if err != nil {
		renderJSON(rw, http.StatusBadRequest, map[string]string{"message": "Invalid request body."})
		return
	}

	renderJSON(rw, http.StatusCreated, product)
}

// GET "/products/{id}/"
func GetProductHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	p := Product{
		Id:    vars["id"],
		Title: "Helloworld",
	}

	renderJSON(rw, http.StatusOK, p)
}

// PUT "/products/{id}/"
func UpdateProductHandler(rw http.ResponseWriter, req *http.Request) {
	oldProduct := Product{
		Id:    "123",
		Title: "Helloworld",
	}
	vars := mux.Vars(req)

	if vars["id"] != oldProduct.Id {
		// let's say id not found.
		renderJSON(rw, http.StatusBadRequest, map[string]string{"message": fmt.Sprintf("Product #%s is not found", vars["id"])})
		return
	}

	// decode new Product data
	var newProduct Product
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		renderJSON(rw, http.StatusBadRequest, map[string]string{"message": "Invalid request body."})
		return
	}

	// update oldProduct
	oldProduct.Title = newProduct.Title
	renderJSON(rw, http.StatusOK, oldProduct)
}

// DELETE "/products/{id}/"
func DeleteProductHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	response := map[string]string{
		"message": fmt.Sprintf("Product #%s has succesfully been deleted", vars["id"]),
	}
	renderJSON(rw, http.StatusOK, response)
}

// JSON RENDERER
func renderJSON(rw http.ResponseWriter, httpStatus int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		response, _ = json.Marshal(map[string]string{"message": "Internal server error."})
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(httpStatus)
	rw.Write(response)
}
