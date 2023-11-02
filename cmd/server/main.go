package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/db"
	"github.com/sushiAlii/salsila/pkg/routes"
)

func main() {
	fmt.Println("Server Initializing...")
	
	db.DB = db.InitializeDB()
	port := "6000"

	r := mux.NewRouter()

	fmt.Printf("Server is running on Port %s", port)

	routes.ConfigureRoleRoutes(r)
	routes.ConfigureSocialNetworkRoutes(r)
	routes.ConfigureFamilyRoutes(r)
	routes.ConfigureAuthRoutes(r)
	routes.ConfigureUserRoutes(r)
	
	err := http.ListenAndServe(":" + port, r)

	if err != nil {
		log.Fatalf("Server failed to start due to error: %v", err)
	}


}