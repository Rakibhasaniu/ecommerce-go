package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"ecommerec/config"
	"ecommerec/rest/middleware"
)

func Start(cnf config.Config) {
	manager := middleware.NewManager()
	mux := http.NewServeMux()
	wrappedMux := manager.With(mux, middleware.Logger, middleware.Cors, middleware.Preflight)
	initRoutes(mux, manager)

	addr := ":" + strconv.Itoa(int(cnf.HttpPort))

	fmt.Print("Starting server on port", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Print("Error starting server:", err)
		os.Exit(1)
	}
}
