package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required,min=3,max=20"`
}

func main() {
	r := gin.Default()
	r.POST("user", userHandler)
	r.Run()
}

func userHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Code": http.StatusOK,
			"Msg":  err.Error(),
		})

		return
	}

	id := user.ID
	name := user.Name
	c.JSON(http.StatusOK, gin.H{
		"Code": http.StatusOK,
		"id":   id,
		"name": name,
	})
}
