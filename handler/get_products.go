package handler

// import "fmt"

import (
	"net/http"

	"ecommerec/database"
	"ecommerec/util"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {

	// if r.Method == "OPTIONS" {
	// 	w.WriteHeader(200)
	// 	return
	// }
	util.SendData(w, database.ProductList, 200)
}