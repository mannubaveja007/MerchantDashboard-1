package main

import (
	"log"
	"merchant-dashboard/config"
	"merchant-dashboard/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitKafka("localhost:9092") 

	r := gin.Default()
	routes.InitRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
