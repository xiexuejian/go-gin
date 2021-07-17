package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 在不同的域中使用 JSONP 从一个服务器请求数据。如果请求参数中存在 callback，添加 callback 到
// response body 。
func Jsonp(c *gin.Context) {
	data := gin.H{
		"foo": "bar",
	}
	//callback is x
	// 将会输出 : x({\"foo\":\"bar\"})
	c.JSONP(http.StatusOK, data)

}
