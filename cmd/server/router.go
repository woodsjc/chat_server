package main

import (
	"net/http"

    "github.com/woodsjc/chat_server/internal/handlers"
	"github.com/bmizerany/pat"
)

func router() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
    mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	return mux
}
