package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/form_post", postHandler)
	r.POST("/form_array", arrayHandler)
	r.POST("/form_map", mapHandler)

	r.Run()
}

func mapHandler(c *gin.Context) {
	m := c.PostFormMap("user")
	c.JSON(http.StatusOK, gin.H{
		"m": m,
	})
}

func arrayHandler(c *gin.Context) {
	ids := c.PostFormArray("ids")

	c.JSON(http.StatusOK, gin.H{
		"ids": ids,
	})
}

func postHandler(c *gin.Context) {
	message := c.PostForm("message")
	name := c.DefaultPostForm("name", "tutu")
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"name":    name,
	})
}
