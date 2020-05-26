package routers

import (
	"github.com/gorilla/mux"
	"github.com/michaelwp/golang-gorm/controllers"
	"github.com/michaelwp/golang-gorm/middlewares"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/", controllers.Home).Methods("GET")

	auth := router.PathPrefix("/auth").Subrouter()
	auth.Use(middlewares.IsAuthorized)
	auth.HandleFunc("/user", controllers.AddUser).Methods("POST")

	return router
}
