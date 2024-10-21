package main

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"

	"github.com/hmuir28/go-thepapucoin-rest/routes"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8010"
	}

	router := gin.New()

	routes.TransferRoutes(router)

	log.Fatal(router.Run(":" + port))
}
