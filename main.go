package main

import (
    "github.com/gin-gonic/gin"
    "merchant-dashboard/routes"
)

func main() {
    r := gin.Default()

    // Initialize routes
    routes.InitRoutes(r)

    // Start the server
    r.Run(":8080") // Listen on port 8080
}
