package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 中间件 - 全局调用
	//r.Use(loginAuth)
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "login",
		})
	})
	r.POST("register", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "register",
		})
	})

	user := r.Group("user", loginAuth)
	{
		// 需要登录保护
		user.GET(":id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "获取用户信息，需要登录保护",
			})
		})

		// 需要登录保护
		user.PUT(":id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "修改用户信息，需要登录保护",
			})
		})
	}

	r.Run()

}

func loginAuth(c *gin.Context) {
	fmt.Println("我是登录保护中间件")
}
