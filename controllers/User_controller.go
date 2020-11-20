package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"login-register/config"
	"login-register/config/auth"
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
	user.Persiapan("tambah")
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

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// decode data json request ke struct user
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	// log.Print(user)
	user.Persiapan("")
	err = user.Validasi("Login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := Proses_login(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func Proses_login(email string, password string) (map[string]string, error) {

	var err error

	user := models.User{}

	err = config.Database.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		// log.Print(err)
		return nil, err
	}
	match, err := models.VerifikasiPassword(user.Password, password)
	if match == false {
		return nil, err
	}
	return auth.BuatToken(user.ID)
}

func Home(w http.ResponseWriter, r *http.Request) {
	tester := "holaaaaa"
	responses.JSON(w, http.StatusOK, tester)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	type token_refresh struct {
		Refresh_token string `json:"refresh_token"`
	}
	var token_old token_refresh
	var new_token map[string]string
	err := json.NewDecoder(r.Body).Decode(&token_old)
	if err != nil {
		panic(err.Error())
	}
	if auth.TokenCek(token_old.Refresh_token) != nil {
		responses.JSON(w, http.StatusInternalServerError, "Expired")
		return
	} else {
		id, err := auth.ExtractTokenID(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		new_token, err = auth.BuatToken(id)
		if err != nil {
			formattedError := formaterror.FormatError(err.Error())
			responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
			return
		}
		responses.JSON(w, http.StatusOK, new_token)
	}

}
