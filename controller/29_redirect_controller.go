package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//	发布HTTP重定向很容易，支持内部和外部链接
func Redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
}

//	从POST发出HTTP重定向。
func PostRedirect(c *gin.Context) {
	c.Redirect(http.StatusFound, "/foo")
}

func TestRedirect(c *gin.Context) {
	c.JSON(200, gin.H{"hello": "world"})
}
