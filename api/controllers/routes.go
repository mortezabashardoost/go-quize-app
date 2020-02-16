package controllers

import (
	"github.com/mortezabashardoost/go-quize-app/api/middlewares"
)

func (server *Server) InitializeRoutes() {
	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Quize Routes
	s.Router.HandleFunc("/quizes", middlewares.SetMiddlewareJSON(s.CreateQuize)).Methods("POST")
}