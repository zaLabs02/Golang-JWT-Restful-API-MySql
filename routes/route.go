package routes

import (
	"login-register/controllers"
	"login-register/middlewares"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/refresh", middlewares.RenderKeJSON(middlewares.SetMiddlewareAuthentication(controllers.Refresh))).Methods("POST")
	router.HandleFunc("/auth/login", middlewares.RenderKeJSON(controllers.Login)).Methods("POST")

	router.HandleFunc("/api/home", middlewares.RenderKeJSON(middlewares.SetMiddlewareAuthentication(controllers.Home))).Methods("GET")

	router.HandleFunc("/api/users", middlewares.RenderKeJSON(controllers.GetSemuaUsers)).Methods("GET")
	router.HandleFunc("/api/users/{id}", middlewares.RenderKeJSON(controllers.GetUserById)).Methods("GET")
	router.HandleFunc("/api/users/{id}", middlewares.RenderKeJSON(middlewares.SetMiddlewareAuthentication(controllers.UpdateUser))).Methods("PUT")
	router.HandleFunc("/api/users", middlewares.RenderKeJSON(controllers.TambahUser)).Methods("POST")
	return router
}
