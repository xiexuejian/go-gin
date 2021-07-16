package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FormPost(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "xxj")
	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}
