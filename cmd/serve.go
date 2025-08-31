package cmd

import (
	"ecommerec/global_router"
	"ecommerec/handler"
	"ecommerec/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	mux := http.NewServeMux()
	// m := manager.With(middleware.Logger)
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handler.GetProductsHandler), middleware.Logger))
	// mux.Handle("GET /products", corsMiddleware(http.HandlerFunc(getProductsHandler)))
	mux.Handle("POST /products", http.HandlerFunc(handler.CreateProduct))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handler.GetProductById))

	fmt.Print("Starting server on :8000\n")

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":8000", globalRouter)
	if err != nil {
		fmt.Print("Error starting server:", err)
	}
}
