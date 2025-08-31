package global_router

import (
	"net/http"
)

func GlobalRouter(mux *http.ServeMux) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,UPDATE,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type") 
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		} else {
			mux.ServeHTTP(w, r)
		}

	}

	return http.HandlerFunc(handleAllReq)

}