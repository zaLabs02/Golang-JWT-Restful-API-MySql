package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"login-register/config"
	"login-register/config/auth"
	"login-register/models"
	"login-register/responses"
	"login-register/responses/formaterror"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetSemuaUsers(w http.ResponseWriter, r *http.Request) {

	user := models.User{}
	users, err := user.ListSemuaUsers(config.Database)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// konversi id dari tring ke int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}
	// log.Print(id)

	user := models.User{}
	userGotten, err := user.LihatUser(config.Database, uint32(id))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// konversi id dari tring ke int
	id, errr := strconv.Atoi(params["id"])

	if errr != nil {
		responses.ERROR(w, http.StatusBadRequest, errr)
		return
	}
	// log.Print(id)

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != uint32(id) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	user.Persiapan("update")
	err = user.Validasi("update")

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedUser, err := user.UpdateDataUser(config.Database, uint32(id))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedUser)
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

func HapusData(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	// konversi id dari tring ke int
	id, errr := strconv.Atoi(params["id"])

	if errr != nil {
		responses.ERROR(w, http.StatusBadRequest, errr)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != 0 && tokenID != uint32(id) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	_, err = user.DeleteAUser(config.Database, uint32(id))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", id))
	responses.JSON(w, http.StatusNoContent, "Data sukses terhapus")

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
