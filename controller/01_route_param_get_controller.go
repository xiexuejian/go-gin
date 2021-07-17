package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// 路由参数（路径占位符）
/**
在对应路径上添加“:+名称”
*/
func TestGet(c *gin.Context) {
	param := c.Param("name")
	log.Infoln(param)
	c.String(200, param)
}

func TestGet02(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

func TestGet03(c *gin.Context) {
	b := c.FullPath() == "/user/:name/action"
	c.String(http.StatusOK, strconv.FormatBool(b))
}
