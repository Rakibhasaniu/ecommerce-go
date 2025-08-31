package handler

// import "fmt"

import (
	"encoding/json"
	"net/http"

	"ecommerec/database"
	"ecommerec/util"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	// if r.Method == "OPTIONS" {
	// 	w.WriteHeader(200)
	// 	return
	// }
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)
	util.SendData(w, newProduct, 201)
}