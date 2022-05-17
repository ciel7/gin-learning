package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 获取所有的文章信息
	r.GET("/posts", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "get",
		})
	})

	// 生成一篇新的文章
	r.POST("/posts", func(c *gin.Context) {
		c.String(http.StatusOK, "post")
	})

	// 修改一篇文章
	r.PUT("/posts/:id", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("PUT id: %s", c.Param("id")))
	})

	// 删除一篇文章
	r.DELETE("/posts/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "delete")
	})

	// 匹配任何方法
	r.Any("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "users any")
	})

	r.Run()
}
