package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 非泛绑定
	r.GET("/user/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")

		c.JSON(http.StatusOK, gin.H{
			"name":   name,
			"action": action,
		})
	})

	// 泛绑定
	r.GET("/posts/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")

		c.JSON(http.StatusOK, gin.H{
			"name":   name,
			"action": action,
		})
	})

	r.Run()
}
