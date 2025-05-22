package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

var users = []User{
    {ID: "1", Name: "Alice"},
    {ID: "2", Name: "Bob"},
}

func main() {
	r := gin.Default()

	// REST endpoints
	r.GET("/hello", helloWorld)
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUserByID)
	r.POST("/users", createUser)

	// Start the server on port 8080
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

// Function for hello world endpoint
func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

// Function to get all users
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// Function to get a user by ID
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

// Function to create a new user
func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}