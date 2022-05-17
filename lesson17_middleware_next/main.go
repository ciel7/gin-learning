package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.Use(middleware1)
	r.GET("ping", func(c *gin.Context) {
		log.Println("function in ...")
		k := c.GetInt("key")
		c.Set("key", k+2000)
		log.Println(k)
		log.Println("function done ...")
	})
	r.Run()
}

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
