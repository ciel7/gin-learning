package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Person struct {
	Name       string    `form:"name"`
	Address    string    `form:"address"`
	Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}

func main() {
	r := gin.Default()
	r.POST("testing", TestHandler)
	r.Run()
}

func TestHandler(c *gin.Context) {
	var p Person
	//ShouldBind checks the Content-Type to select a binding engine automatically
	// 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
	// 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）。
	// 查看更多：https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if err := c.ShouldBind(&p); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"Code":       0,
			"Name":       p.Name,
			"Address":    p.Address,
			"Birthday":   p.Birthday,
			"UnixTime":   p.UnixTime,
			"CreateTime": p.CreateTime,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Code": http.StatusOK,
			"Msg":  err.Error(),
		})
	}

	return
}
