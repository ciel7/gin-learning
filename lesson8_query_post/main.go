package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID string `form:"id" binding:"required,uuid"`
}

func main() {
	r := gin.Default()

	r.GET("user", func(c *gin.Context) {
		var user User

		//if err := c.ShouldBindQuery(&user); err != nil {
		if err := c.BindQuery(&user); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Code": http.StatusBadRequest,
				"Msg":  err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Code": 0,
			"Id":   user.ID,
		})
	})
	r.Run()
}
