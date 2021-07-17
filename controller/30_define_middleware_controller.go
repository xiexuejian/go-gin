package controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

func SelfDefineMid(c *gin.Context) {
	example := c.MustGet("example").(string)
	// 它将打印: "12345"
	log.Println(example)
}
