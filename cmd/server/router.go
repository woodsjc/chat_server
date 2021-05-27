package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/woodsjc/chat_server/internal/handlers"
)

func router() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	return mux
}
