package http

import (
	"log"
	"net/http"
	"poke-store/internal/adapters/http/router"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer() Server {
	router := mux.NewRouter().StrictSlash(true)
	return Server{router: router}
}

func (s *Server) Start() {
	productRouter := router.NewProductRouter(s.router)
	productRouter.ConfigureRoutes()
	log.Fatal(http.ListenAndServe(":8080", s.router))
}
