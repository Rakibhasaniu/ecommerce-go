package handler

import (
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

	product := database.Get(pid)
	if product == nil {
		util.SendError(w, 404, "Product Not Found")
		return
	}

	util.SendData(w, product, 200)

}
