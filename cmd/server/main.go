package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sushiAlii/salsila/pkg/controllers"
	"github.com/sushiAlii/salsila/pkg/db"
	"github.com/sushiAlii/salsila/pkg/models"
	"github.com/sushiAlii/salsila/pkg/routes"
)

func main() {
	fmt.Println("Server Initializing...")
	
	dbInstance := db.InitializeDB()
	port := os.Getenv("APP_PORT")

	r := mux.NewRouter()
	s := models.InstantiateServices(dbInstance)
	c := controllers.InstantiateControllers(s)
	
	routes.ConfigureAllRoutes(r, c)
	
	err := http.ListenAndServe(":" + port, r)

	if err != nil {
		log.Fatalf("Server failed to start due to error: %v", err)
	}
}