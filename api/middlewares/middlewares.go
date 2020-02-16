package middlewares

import (
	"errors"
	"net/http"

	"github.com/mortezabashardoost/go-quize-app/api/responses"
)

func SetMiddlewareJSON(next http.HandleFunc) http.HandleFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type":"application/json")
		next(w,r)
	}
}