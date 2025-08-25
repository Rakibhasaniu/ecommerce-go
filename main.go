package main

import "fmt"

import "net/http"


func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello, world!")
}
func aboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "About page")
}

func main(){
	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Hello, world!")
	// })

	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)

	fmt.Print("Starting server on :3000\n")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Print("Error starting server:", err)
	}
	// http.ListenAndServe(":3000", mux)
}