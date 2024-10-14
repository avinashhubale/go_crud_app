package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/avinashhubale/go-crud-app/handlers"
)


func SetupRoutes(r *gin.Engine) {
    r.GET("/users", handlers.GetUsers)
    r.GET("/users/:id", handlers.GetUserByID)
    r.POST("/users", handlers.CreateUser)
    r.PUT("/users/:id", handlers.UpdateUser)
    r.DELETE("/users/:id", handlers.DeleteUser)
}
