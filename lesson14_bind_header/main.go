package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Header struct {
	ID string `header:"id" binding:"required"`
}

func main() {
	r := gin.Default()
	// http://localhost:8080/user/1234
	r.POST("test", TestHandler)
	r.Run()
}

func TestHandler(c *gin.Context) {
	var h Header
	if err := c.ShouldBindHeader(&h); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"Code": 0,
			"ID":   h.ID,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Code": http.StatusOK,
			"Msg":  err.Error(),
		})
	}
}
