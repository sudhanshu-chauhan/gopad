package routes

import (
	"github.com/gorilla/mux"

	"github.com/sudhanshu-chauhan/gopad/controllers"
)

func GetIndexRoutes() mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/person", controllers.PostPerson).Methods("POST")
	return *router
}
