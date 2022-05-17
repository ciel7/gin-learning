package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID       string `form:"id" binding:"required,uuid"`
	UserName string `form:"user_name" binding:"required,min=3,max=8"`
}

func main() {
	r := gin.Default()
	r.POST("user", UserHandler)
	r.Run()
}

func UserHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Code": http.StatusOK,
			"Msg":  err.Error(),
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Code":     0,
		"ID":       user.ID,
		"UserName": user.UserName,
	})
}
