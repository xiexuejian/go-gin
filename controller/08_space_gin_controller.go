package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * 第一种：r := gin.New()---->不带中间件

 * 第二种：默认已经连接了 Logger and Recovery 中间件
 * r := gin.Default()
 */
func SpaceGin(c *gin.Context) {
	c.String(http.StatusOK, "执行成功")
}
