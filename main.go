package main

import (
	"go_web/config"
	"go_web/controllers/categorycontroller"
	"go_web/controllers/homecontroller"
	"go_web/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDb()
	// memanggil homepage
	http.HandleFunc("/" , homecontroller.Welcome)

	// categories
	http.HandleFunc("/categories" , categorycontroller.Index)
	http.HandleFunc("/categories/add" , categorycontroller.Add)
	http.HandleFunc("/categories/edit" , categorycontroller.Edit)
	http.HandleFunc("/categories/delete" , categorycontroller.Delete)

	// product
	http.HandleFunc("/products" , productcontroller.Index)		
	http.HandleFunc("/products/detail" , productcontroller.Detail)		
	http.HandleFunc("/products/create" , productcontroller.Add)		
	http.HandleFunc("/products/update" , productcontroller.Update)		
	http.HandleFunc("/products/delete" , productcontroller.Delete)		

	log.Println("server running at port 8000")
	// membutuhkan 2 param address / port , handler
	http.ListenAndServe(":8000",nil)
}
