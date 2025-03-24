package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"file-storage-api/app/routes"
	"file-storage-api/app/utils"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func main() {

	utils.InitLogger()
	utils.Logger.Info("Starting the application...")

	utils.LoadEnv()

	client = utils.InitDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create Router
	router := mux.NewRouter()

	routes.RegisterRoutes(router)

	// Test Endpoint
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API is running...")
	})

	// Configure CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Wrap router with CORS handler
	handler := c.Handler(router)

	fmt.Println("🚀 Server is running on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
