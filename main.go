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
	// r.GET("/users", getUsers)
	// r.GET("/users/:id", getUserByID)
	// r.POST("/users", createUser)

	// Start the server on port 8080
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}