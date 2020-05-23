package main

import (
	"fmt"
	"golang-gorm/routers"
	"log"
	"net/http"
)

func main() {
	port := ":8000"
	router := routers.Router()

	fmt.Println("server listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
