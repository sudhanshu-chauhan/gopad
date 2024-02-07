package routes

import (
	"github.com/gorilla/mux"

	"gopad/controllers"
)

func GetIndexRoutes() mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/person", controllers.PostPerson).Methods("POST")
	return *router
}
