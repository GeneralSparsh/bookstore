package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/priya/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter() //initialising router
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r)) //helps create the server and the parameter is address on which we want to start the server
}
