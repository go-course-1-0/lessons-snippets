package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		//panic("errorrrrr!")

		c.JSON(http.StatusOK, gin.H{
			"message": "Hello!",
		})
	})

	r.GET("/users", GetUsers)

	r.Run(":6767")
}

func GetUsers(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusAlreadyReported, gin.H{
		"message": "I'm on vacation!",
	})
}
