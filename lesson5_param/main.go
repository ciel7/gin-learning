package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Age  int    `uri:"id"`
	Name string `uri:"name"`
}

func main() {
	r := gin.Default()

	//r.GET("/posts", func(c *gin.Context) {
	//	c.String(http.StatusOK, "1")
	//})

	//r.GET("/:id/:name", func(c *gin.Context) {
	//	id := c.Param("id")
	//	name := c.Param("name")
	//	c.JSON(http.StatusOK, gin.H{
	//		"id":   id,
	//		"name": name,
	//	})
	//})

	r.GET("/:id/:name", func(c *gin.Context) {
		var p Person
		// ShouldBindUri binds the passed struct pointer using the specified binding engine.
		if err := c.ShouldBindUri(&p); err != nil {
			// strconv.ParseInt: parsing "tutu": invalid syntax
			fmt.Println("err = ", err)
			c.Status(404)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"name": p.Name,
			"age":  p.Age,
		})
	})

	r.Run()
}
