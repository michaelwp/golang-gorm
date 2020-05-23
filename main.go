package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":8000"
	fmt.Println("server listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
