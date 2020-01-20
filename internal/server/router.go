package server

import ("github.com/gorilla/mux"
"github.com/natalie-todd/anagrams_go/internal/server/handlers")

type Router struct {
	*mux.Router
	handler *handlers.Handler
}