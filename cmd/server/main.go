package main

import (
	"log"
	"net/http"

	"github.com/woodsjc/chat_server/internal/handlers"
)

func main() {
	mux := router()

	go handlers.ListenToWsChannel()
	log.Fatal(http.ListenAndServe(":8080", mux))
}
