package handler

import (
	"log"
	"net/http"
	"strconv"

	"ecommerec/database"
	"ecommerec/util"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	pid, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please provide a valid id", 400)
		return
	}

	for idx, product := range database.ProductList {
		log.Println(idx, product)

		if product.ID == pid {
			util.SendData(w, product, 200)
			return
		}
	}
	util.SendData(w, "Data pai nai halarput ki id dis baintost", 404)

}
