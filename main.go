package main

import (
	"fmt"
	"golang-gorm/db"
	"golang-gorm/routers"
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
