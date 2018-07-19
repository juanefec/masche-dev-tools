package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"bytes"
	"encoding/json"
	"io/ioutil"
)

func main() {
	router := gin.Default()
	// Serve frontend static files
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
	body := getToken()
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"token": body,
	})
}

// LikeJoke increments the likes of a particular joke Item
func LikeJoke(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "LikeJoke handler not implemented yet",
	})
}

type Token struct {
	Name   string
	Token  string
	UserID string
}

func getToken() string {
	url := "http://webcalldesa02:9300/api/tokens"

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
	return token.Token
}
