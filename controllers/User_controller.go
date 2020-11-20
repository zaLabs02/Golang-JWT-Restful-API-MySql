package controllers

import (
	"encoding/json"
	"fmt"
	"login-register/config"
	"login-register/models"
	"login-register/responses"
	"login-register/responses/formaterror"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	user := models.User{}
	users, err := user.ListSemuaUsers(config.Database)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func TambahUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// decode data json request ke struct user
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	// log.Print(user)
	user.Persiapan()
	err = user.Validasi("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.TmbhUser(config.Database)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, userCreated)
}
