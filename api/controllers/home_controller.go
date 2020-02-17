package controllers

import (
	"net/http"

	"go-quize-app/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to Quize App API")
}
