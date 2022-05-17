package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterForm struct {
	UserName string `json:"username" binding:"required,min=3,max=7"`
	Password string `json:"password" binding:"required,len=8"`
	Age      uint32 `json:"age" binding:"required,gte=1,lte=150"`
	Sex      uint32 `json:"sex" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginForm struct {
	UserName   string `json:"username" binding:"required,min=3,max=7"`
	Password   string `json:"password" binding:"required,len=8"`
	RePassword string `json:"re_password" binding:"required,len=8"`
}

func main() {
	r := gin.Default()
	r.POST("/login", LoginHandler)
	r.POST("/register", RegisterHandler)
	r.Run()
}

// RegisterHandler 注册
func RegisterHandler(c *gin.Context) {
	var r RegisterForm
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 40002,
			"msg":  "注册失败，请检查参数",
			"err":  err.Error(),
		})
		return
	}

	// 注册成功
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "注册成功",
		"data": r,
	})
}

// LoginHandler 登录
func LoginHandler(c *gin.Context) {
	var l LoginForm
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 40001,
			"msg":  "登录失败，请检查参数",
			"err":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登陆成功",
		"data": l.UserName,
	})
}
