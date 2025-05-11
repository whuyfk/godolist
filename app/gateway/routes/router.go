package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/whuyfk/godolist/app/gateway/internal/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	RegisterRouter(router)
	return router
}

func RegisterRouter(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	// 任务模块
	task := router.Group("/task")
	{
		task.GET("/hello", http.SayHello)
		task.POST("/create", http.TaskCreate)
	}
}
