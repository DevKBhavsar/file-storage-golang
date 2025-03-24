package routes

import (
	"github.com/gorilla/mux"
)

// RegisterRoutes initializes all routes
func RegisterRoutes(router *mux.Router) {
	AuthRoutes(router)
	FileRoutes(router)
}
