package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.New()
	// 使用 Logger 中间件
	r.Use(gin.Logger())
	// 使用 Recover 中间件
	r.Use(gin.Recovery())
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
