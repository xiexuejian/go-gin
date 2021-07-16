package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestGet04(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	// 这个是c.Request.URL.Query().Get("lastname") 快捷写法
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)

}
