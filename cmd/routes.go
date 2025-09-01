package cmd

import (
	"net/http"

	"ecommerec/handler"
	"ecommerec/middleware"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handler.GetProductsHandler)))
	// mux.Handle("GET /products", corsMiddleware(http.HandlerFunc(getProductsHandler)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handler.CreateProduct)))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handler.GetProductById)))

}
