package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/avinashhubale/go-crud-app/database"
    "github.com/avinashhubale/go-crud-app/models"
)

// GetUsers retrieves all users from the database
func GetUsers(c *gin.Context) {
    var users []models.User
    err := database.DB.Select(&users, "SELECT * FROM users")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

// GetUserByID retrieves a user by ID
func GetUserByID(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    err := database.DB.Get(&user, "SELECT * FROM users WHERE id=$1", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := database.DB.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

// UpdateUser updates an existing user
func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := database.DB.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", user.Name, user.Email, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

// DeleteUser deletes a user by ID
func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    _, err := database.DB.Exec("DELETE FROM users WHERE id=$1", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
