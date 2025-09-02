package rest

import (
	"net/http"

	"ecommerec/rest/handler"
	"ecommerec/rest/middleware"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handler.GetProductsHandler)))
	// mux.Handle("GET /products", corsMiddleware(http.HandlerFunc(getProductsHandler)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handler.CreateProduct)))
	mux.Handle("GET /product/{id}", manager.With(http.HandlerFunc(handler.GetProductById)))
	mux.Handle("PUT /product/{id}", manager.With(http.HandlerFunc(handler.UpdateProduct)))
	mux.Handle("DELETE /product/{id}", manager.With(http.HandlerFunc(handler.DeleteProduct)))

}
