package routes

import (
	"file-storage-api/app/controllers"
	"github.com/gorilla/mux"
)

// AuthRoutes defines user authentication routes
func AuthRoutes(router *mux.Router) {
	router.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", controllers.LoginUser).Methods("POST")

}
