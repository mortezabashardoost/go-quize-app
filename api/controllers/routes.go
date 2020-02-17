package controllers

import (
	"go-quize-app/api/middlewares"
)

func (s *Server) initializeRoutes() {
	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Quize Routes
	s.Router.HandleFunc("/quizes", middlewares.SetMiddlewareJSON(s.CreateQuize)).Methods("POST")
}
