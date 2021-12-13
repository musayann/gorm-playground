package main

import (
	"log"
	config "musayann/gorm-playground/config"
	"musayann/gorm-playground/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}
