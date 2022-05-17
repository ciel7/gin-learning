package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 构建一个自带默认中间件的 *engine
	r := gin.Default()
	//r.Use(middleware1)
	r.Use(RefererMiddleware())
	r.Use(func(c *gin.Context) {
		log.Println("我是另一个中间件 ... ")
	})
	r.GET("ping", func(c *gin.Context) {
		//log.Println("function in ...")
		//k := c.GetInt("key")
		//c.Set("key", k+2000)
		//log.Println(k)
		//log.Println("function done ...")
		c.JSON(200, gin.H{
			"code": 0,
		})
	})
	r.Run()
}

// 中间件第一种实现方式
func middleware1(c *gin.Context) {
	log.Println("in middleware1 ...")
	c.Set("key", 1000)
	log.Println("before next ...")
	c.Next()
	log.Println("after next ...")
	log.Println("middleware1 done...")
	c.JSON(200, gin.H{
		"msg": c.GetInt("key"),
	})
}

// 中间件第二种实现方式
// Gin的中间件，实际上就是 Gin 定义的一个 HandlerFunc
func middleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": c.GetInt("key"),
		})
	}
}

// RefererMiddleware Header Referer 中间件
func RefererMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 取到 referer
		ref := c.GetHeader("Referer")

		if ref == "" {
			c.AbortWithStatusJSON(200, gin.H{
				"msg": "非法访问",
			})
			return
		}
		c.Next()
	}
}
