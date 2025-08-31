package cmd

import (
	"fmt"
	"net/http"

	"ecommerec/global_router"
	"ecommerec/handler"
)


func Serve (){
	mux := http.NewServeMux()
	mux.Handle("GET /products", http.HandlerFunc(handler.GetProductsHandler))
	// mux.Handle("GET /products", corsMiddleware(http.HandlerFunc(getProductsHandler)))
	mux.Handle("POST /products", http.HandlerFunc(handler.CreateProduct))
	mux.Handle("GET /products/{id}",http.HandlerFunc(handler.GetProductById))

	fmt.Print("Starting server on :8000\n")

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":8000", globalRouter)
	if err != nil {
		fmt.Print("Error starting server:", err)
	}
}