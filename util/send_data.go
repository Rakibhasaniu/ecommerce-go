package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	// json.NewEncoder(w).Encode(newProduct)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
