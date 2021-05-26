package main

import (
	"log"
	"net/http"
)

func main() {
	mux := router()

	log.Fatal(http.ListenAndServe(":8080", mux))
}
