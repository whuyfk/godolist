package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	app.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello",
		})
	})
	app.Run()
}
