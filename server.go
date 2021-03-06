package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Token struct {
	Name   string
	Token  string
	UserID string
}

type Login struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type ChangePassword struct {
	User        string `json:"user"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

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
	router.Use(CORSMiddleware())
	router.Use(static.Serve("/", static.LocalFile("./react-client/build", true)))

	api := router.Group("/api")
	{
		api.POST("/token", getToken)
		api.POST("/changePassword", changePassword)
	}

	router.Run(":3000")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3002")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func getToken(c *gin.Context) {
	var login Login
	if err := c.ShouldBind(&login); err == nil {
		url := "http://webcalldesa02:9300/api/tokens"
		var jsonStr = []byte(`{"username":"` + login.User + `","password":"` + login.Password + `"}`)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
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
		c.JSON(http.StatusOK, gin.H{"token": token.Token})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func changePassword(c *gin.Context) {
	var request ChangePassword
	if err := c.ShouldBind(&request); err == nil {
		url := "http://webcrmdesaf03:8600/api/LoginData/AlterPassword?user=" + request.User + "&oldPassword=" + request.OldPassword + "&newPassword=" + request.NewPassword
		req, _ := http.NewRequest("GET", url, nil)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusInternalServerError {
			c.JSON(http.StatusInternalServerError, "No se pudo cambiar la contraseña")
		} else {
			c.JSON(http.StatusNoContent, nil)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
