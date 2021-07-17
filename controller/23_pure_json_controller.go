package controller

import "github.com/gin-gonic/gin"

// 通常情况下，JSON会将特殊的HTML字符替换为对应的unicode字符，比如<替换为\u003c，如果想原
// 样输出html，则使用PureJSON，这个特性在Go 1.6及以下版本中无法使用。

func JsonOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"html": "<b>Hello, world!</b>",
	})
}

func PureJSON(c *gin.Context) {
	c.PureJSON(200, gin.H{
		"html": "<b>Hello, world!</b>",
	})
}
