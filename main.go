package main

import (
	"ecommerec/cmd"
)

// func corsMiddleware(next http.Handler) http.Handler {
// 	handleCors := func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,UPDATE,PATCH,DELETE,OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 		next.ServeHTTP(w, r)
// 	}

// 	handler := http.HandlerFunc(handleCors)
// 	return handler
// }


func main() {
	cmd.Serve()
	// http.ListenAndServe(":3000", mux)
}


