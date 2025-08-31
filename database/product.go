package database


var ProductList []Product

type Product struct {
	ID          int
	Title       string
	Price       float64
	Description string
	ImgUrl      string
}



func init() {
	prd1 := Product{ID: 1, Title: "Product 1", Price: 10.0, Description: "Description for database 1", ImgUrl: "https://via.placeholder.com/150"}
	prd2 := Product{ID: 2, Title: "Product 2", Price: 20.0, Description: "Description for database 2", ImgUrl: "https://via.placeholder.com/150"}
	prd3 := Product{ID: 3, Title: "Product 3", Price: 30.0, Description: "Description for database 3", ImgUrl: "https://via.placeholder.com/150"}
	prd4 := Product{ID: 4, Title: "Product 4", Price: 40.0, Description: "Description for database 4", ImgUrl: "https://via.placeholder.com/150"}
	prd5 := Product{ID: 5, Title: "Product 5", Price: 50.0, Description: "Description for database 5", ImgUrl: "https://via.placeholder.com/150"}
	prd6 := Product{ID: 6, Title: "Product 6", Price: 60.0, Description: "Description for database 6", ImgUrl: "https://via.placeholder.com/150"}

	ProductList = append(ProductList, prd1, prd2, prd3, prd4, prd5, prd6)
}
