package controllers

import (
	"net/http"

	"github.com/mortezabashardoost/go-quize-app/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOk, "Welcome to Quize App API")
}