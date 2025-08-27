package main

import "fmt"

import "net/http"
import "encoding/json"


func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello, world!")
}
func aboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "About page")
}

type Product struct {
	ID int `json:"rakib_id"`
	Title string
	Price float64
	Description string
	ImgUrl string
}

var productList  []Product

func getProductsHandler( w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	encoder :=json.NewEncoder(w)
	err := encoder.Encode(productList)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func createProduct (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newProduct  Product

	decoder :=json.NewDecoder(r.Body)
	err :=decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/products", getProductsHandler)

	fmt.Print("Starting server on :3000\n")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Print("Error starting server:", err)
	}
	// http.ListenAndServe(":3000", mux)
}

func init(){
	prd1 := Product{ID: 1, Title: "Product 1", Price: 10.0, Description: "Description for product 1", ImgUrl: "https://via.placeholder.com/150"}	
	prd2 := Product{ID: 2, Title: "Product 2", Price: 20.0, Description: "Description for product 2", ImgUrl: "https://via.placeholder.com/150"}
	prd3 := Product{ID: 3, Title: "Product 3", Price: 30.0, Description: "Description for product 3", ImgUrl: "https://via.placeholder.com/150"}
	prd4 := Product{ID: 4, Title: "Product 4", Price: 40.0, Description: "Description for product 4", ImgUrl: "https://via.placeholder.com/150"}
	prd5 := Product{ID: 5, Title: "Product 5", Price: 50.0, Description: "Description for product 5", ImgUrl: "https://via.placeholder.com/150"}
	prd6 := Product{ID: 6, Title: "Product 6", Price: 60.0, Description: "Description for product 6", ImgUrl: "https://via.placeholder.com/150"}

	productList = append(productList, prd1, prd2, prd3, prd4, prd5, prd6)
}