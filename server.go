package main

import (
	"net/http"

	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func datamock() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/tokens", func(c *gin.Context) {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusOK, gin.H{
				"token":  "holajajasoyuntoken",
				"name":   "Pepito",
				"userID": "gato",
			})
		})
	}
	router.Run(":3001")
}

func main() {
	go datamock()
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(static.Serve("/", static.LocalFile("./react-client/build", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/Token", TokenHandler)
	}

	router.POST("/jokeLike", LikeJoke)

	// Start and run the server
	router.Run(":3000")
}

// TokenHandler returns a token as response
func TokenHandler(c *gin.Context) {
	token, userID := getToken()
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"token":  token,
		"userId": userID,
	})
}

// LikeJoke increments the likes of a particular joke Item
func LikeJoke(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "LikeJoke handler not implemented yet",
	})
}

// Token is the structure that the service resturns
type Token struct {
	Name   string
	Token  string
	UserID string
}

func getToken() (string, string) {
	url := "http://localhost:3001/api/tokens"

	var jsonStr = []byte(`{"username":"B048151","password":"juan1807"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var token Token
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &token)
	return token.Token, token.UserID
}
