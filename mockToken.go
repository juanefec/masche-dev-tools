package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/tokens", TokenHandler)
	}
	router.Run(":3001")
}

// TokenHandler returns a token as response
func TokenHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"token": "holajajasoyuntoken", "name": "Pepito", "userID": "gato"})
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
