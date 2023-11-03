package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/db"
	"github.com/sushiAlii/salsila/pkg/routes"
)

func main() {
	fmt.Println("Server Initializing...")
	
	db.DB = db.InitializeDB()
	port := os.Getenv("APP_PORT")

	r := mux.NewRouter()

	fmt.Printf("Server is running on Port %s", port)

	routes.ConfigureAllRoutes(r)
	
	err := http.ListenAndServe(":" + port, r)

	if err != nil {
		log.Fatalf("Server failed to start due to error: %v", err)
	}
}