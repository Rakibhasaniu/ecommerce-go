package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"ecommerec/database"
	"ecommerec/util"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pid, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
	}
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
	}
	newProduct.ID = pid
	database.Update(newProduct)
	util.SendData(w, "Successfully updated", 201)
}
