package routes

import (
	"login-register/controllers"
	"login-register/middlewares"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/users", middlewares.RenderKeJSON(controllers.GetUsers)).Methods("GET")
	router.HandleFunc("/api/users", middlewares.RenderKeJSON(controllers.TambahUser)).Methods("POST")
	return router
}
