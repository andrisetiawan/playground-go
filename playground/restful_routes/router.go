package main

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
)

func router() *mux.Router {
	r := mux.NewRouter()

	// allow trailing slash
	r.StrictSlash(true)

	// Products
	products := r.PathPrefix("/products").Headers("Content-Type", "application/json").Subrouter()

	// GET "/products/"
	products.HandleFunc("/", GetProductsHandler).Methods("GET")
	// POST "/products"
	products.HandleFunc("/", CreateProductsHandler).Methods("POST")
	// GET "/products/{id}/"
	products.HandleFunc("/{id}/", GetProductHandler).Methods("GET")
	// PUT "/products/{id}/"
	products.HandleFunc("/{id}/", UpdateProductHandler).Methods("PUT")
	// DELETE "/products/{id}/"
	products.HandleFunc("/{id}/", DeleteProductHandler).Methods("DELETE")

	// Page not found handler
	r.NotFoundHandler = http.HandlerFunc(PageNotFoundHandler)

	return r
}

func PageNotFoundHandler(rw http.ResponseWriter, req *http.Request) {
	r := render.New()
	r.JSON(rw, http.StatusOK, map[string]string{"message": "Page not found"})
}
