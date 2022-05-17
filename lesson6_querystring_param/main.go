package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/welcome", welcomeHandler)
	r.GET("/array", arrayHandler)
	r.GET("/map", mapHandler)

	r.Run()
}

func mapHandler(c *gin.Context) {
	// http://127.0.0.1:8080/map?user[name]=tutu&user[age]=18
	m := c.QueryMap("user")

	c.JSON(http.StatusOK, gin.H{
		"m": m,
	})
}

func arrayHandler(c *gin.Context) {
	// http://127.0.0.1:8080/array?ids=1,2,3,4,5
	ids := c.QueryArray("ids")

	c.JSON(http.StatusOK, gin.H{
		"ids": ids,
	})
}

func welcomeHandler(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "tutu")
	lastname := c.Query("lastname")

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("welcome %v %v", lastname, firstname),
	})
}
