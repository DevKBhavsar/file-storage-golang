package routes

import (
	"file-storage-api/app/controllers"
	"file-storage-api/app/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

// FileRoutes defines routes for file management
func FileRoutes(router *mux.Router) {
	router.Handle("/upload", middleware.JWTMiddleware(http.HandlerFunc(controllers.UploadFile))).Methods("POST")
	router.Handle("/files", middleware.JWTMiddleware(http.HandlerFunc(controllers.GetFiles))).Methods("GET")
router.Handle("/storage/remaining", middleware.JWTMiddleware(http.HandlerFunc(controllers.GetRemainingStorage))).Methods("GET")

}
