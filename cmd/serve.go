package cmd

import (
	"fmt"
	"net/http"

	"ecommerec/middleware"
)

func Serve() {
	manager := middleware.NewManager()
	mux := http.NewServeMux()
	fmt.Print("Starting server on :8000\n")
	wrappedMux := manager.With(mux, middleware.Logger, middleware.Cors, middleware.Preflight)
	initRoutes(mux, manager)
	err := http.ListenAndServe(":8000", wrappedMux)
	if err != nil {
		fmt.Print("Error starting server:", err)
	}
}
