package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID       string `form:"id" binding:"required"`
	UserName string `form:"username" binding:"required,min=3"`
	PassWord string `form:"password" binding:"required,min=3"`
}

func main() {
	r := gin.Default()
	r.GET("user", UserHandler)
	r.Run()
}

func UserHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindQuery(&user); err != nil {
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
		"PassWord": user.PassWord,
	})
}
