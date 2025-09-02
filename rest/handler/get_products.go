package handler

// import "fmt"

import (
	"net/http"

	"ecommerec/database"
	"ecommerec/util"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.List(), 200)
}
