package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//	使用 LoadHTMLGlob() 或者 LoadHTMLFiles()
func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}

/**
如果是在不同目录中，我们要区别路径
*/

/**
其它知识点：
1、自定义模板渲染器
2、自定义分隔符
3、自定义模板函数
*/
