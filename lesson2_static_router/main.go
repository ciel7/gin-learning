package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 加载静态文件
	// http://127.0.0.1:8080/images/infoflow%202022-05-02%2013-56-29.png
	r.Static("/images", "./images")

	// StaticFS works just like `Static()` but a custom `http.FileSystem` can be used instead.
	r.StaticFS("/static", http.Dir("./static"))

	// 加载单独的静态文件
	r.StaticFile("index", "index.html")

	r.Run()
}
