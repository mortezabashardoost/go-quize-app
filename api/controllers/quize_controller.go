package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go-quize-app/api/models"
	"go-quize-app/api/responses"
)

func (server *Server) CreateQuize(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	quize := models.Quize{}
	err = json.Unmarshal(body, &quize)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// TODO: implement remaining functions
	// quize.Prepare()
	// err = quize.Validate()
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }
	// uid, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	// if uid != quize.UserId {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
	// 	return
	// }
	quizeCreated, err := quize.SaveQuize(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, quizeCreated.QuizeId))
	responses.JSON(w, http.StatusCreated, quizeCreated)
}
