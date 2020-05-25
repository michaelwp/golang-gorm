package main

import (
	"fmt"
	"github.com/michaelwp/golang-gorm/db"
	"github.com/michaelwp/golang-gorm/routers"
	"log"
	"net/http"
)

func main() {
	port := ":8000"
	router := routers.Router()

	db.MySql()

	fmt.Println("server listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
