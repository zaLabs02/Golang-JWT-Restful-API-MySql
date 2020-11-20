package controllers

import (
	"encoding/json"
	"fmt"
	"login-register/config"
	"login-register/config/auth"
	"login-register/models"
	"login-register/responses"
	"login-register/responses/formaterror"
	"net/http"
	"strconv"
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

func Proses_login(email string, password string) (string, error) {

	var err error

	user := models.User{}

	err = config.Database.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		// log.Print(err)
		return "", err
	}
	match, err := models.VerifikasiPassword(user.Password, password)
	if match == false {
		return strconv.FormatBool(match), err
	}
	return auth.BuatToken(user.ID)
	// log.Print(match, err)
	// return strconv.FormatBool(match), err
}

func Home(w http.ResponseWriter, r *http.Request) {
	tester := "holaaaaa"
	responses.JSON(w, http.StatusOK, tester)
}
