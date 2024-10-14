package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/avinashhubale/go-crud-app/database"
    "github.com/avinashhubale/go-crud-app/routes"
)

func main() {
    // Initialize the database
    database.InitDB()

    // Create a new Gin router
    r := gin.Default()

    // Set up routes
    routes.SetupRoutes(r)

    // Run the server
    if err := r.Run(":8080"); err != nil {
        log.Fatal("Unable to start server: ", err)
    }
}
