package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//p := r.Group("/posts")
	//{
	//	p.GET("/", GetHandler)
	//	p.POST("/posts", PostHandler)
	//	p.PUT("/", PutHandler)
	//	p.DELETE("/:id", DeleteHandler)
	//}

	// localhost:8080/api/v1
	//v1 := r.Group("/v1")
	//{
	//	v1.POST("/login", loginEndPoint)
	//	v1.POST("/submit", submitEndPoint)
	//	v1.POST("/read", readEndPoint)
	//}

	// localhost:8080/api/v2
	//v2 := r.Group("/v2")
	//{
	//	v2.POST("/login", loginEndPoint)
	//	v2.POST("/submit", submitEndPoint)
	//	v2.POST("/read", readEndPoint)
	//}

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/login", loginEndPoint)
			v1.POST("/submit", submitEndPoint)
			v1.POST("/read", readEndPoint)
		}

		v2 := api.Group("/v2")
		{
			v2.POST("/login", loginEndPoint)
			v2.POST("/submit", submitEndPoint)
			v2.POST("/read", readEndPoint)
		}
	}

	r.Run()
}

func readEndPoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "readEndPoint",
	})
}

func submitEndPoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "submitEndPoint",
	})
}

func loginEndPoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "loginEndPoint",
	})
}

func GetHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get",
	})
}

func PostHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "post",
	})
}

func PutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "put",
	})
}

func DeleteHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "delete",
	})
}
