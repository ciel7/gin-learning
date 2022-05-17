package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID string `uri:"id" binding:"required"`
}

func main() {
	r := gin.Default()
	// http://localhost:8080/user/1234
	r.POST("user/:id", UserHandler)
	r.Run()
}

func UserHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindUri(&user); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"Code": 0,
			"ID":   user.ID,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Code": http.StatusOK,
			"Msg":  err.Error(),
		})
	}
}
