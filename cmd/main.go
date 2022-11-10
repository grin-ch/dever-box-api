package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grin-ch/dever-box-api/auth"
)

func main() {
	r := gin.Default()
	r.Use(auth.AuthMiddlewares()...)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
