package routers

import (
	"github.com/gorilla/mux"
	"golang-gorm/controllers"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.Home).Methods("GET")

	return router
}
