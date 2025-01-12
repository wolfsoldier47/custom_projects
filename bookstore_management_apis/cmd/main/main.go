package main

import (
	"fmt"
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	fmt.Println("server starting at 3000")
	log.Fatal(http.ListenAndServe(":3000", r))

}
