package handler

import (
	"net/http"
	"strconv"

	"ecommerec/database"
	"ecommerec/util"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	pid, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
	}

	database.Delete(pid)
	util.SendData(w, "Product Deleted Successfully", 201)

}
