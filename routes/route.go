package routes

import (
	"login-register/controllers"
	"login-register/middlewares"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/auth/login", middlewares.RenderKeJSON(controllers.Login)).Methods("POST")

	router.HandleFunc("/api/home", middlewares.RenderKeJSON(middlewares.SetMiddlewareAuthentication(controllers.Home))).Methods("GET")

	router.HandleFunc("/api/users", middlewares.RenderKeJSON(controllers.GetUsers)).Methods("GET")
	router.HandleFunc("/api/users", middlewares.RenderKeJSON(controllers.TambahUser)).Methods("POST")
	return router
}
