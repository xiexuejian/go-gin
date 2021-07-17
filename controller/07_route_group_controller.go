package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteGroup(c *gin.Context) {
	c.String(http.StatusOK, "执行成功")
}
