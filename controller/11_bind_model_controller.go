package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
若要将请求主体绑定到结构体中，请使用模型绑定，目前支持JSON、XML、YAML和标准表单值
（foo=bar&boo=baz）的绑定。

Gin使用go-playground/validator.v8验证参数，点击此处查看完整文档here
需要在绑定的字段上设置tag，比如，绑定格式为json，需要设置为 json:"fieldname"
此外，Gin提供了两种绑定方法：
	1、类型 - Must bind
	2、类型 - Should bind

*/

//	binding:"required" 	代表该字段不能为空
//	binding:"-"	可以为空
type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func BindJson(c *gin.Context) {
	var json Login

	err := c.ShouldBindJSON(&json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.User != "manu" || json.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

// XML 绑定示例 (
// <?xml version="1.0" encoding="UTF-8"?>
// <root>
// <user>user</user>
// <password>123</password>
// </root>)

func BIndXml(c *gin.Context) {
	var xml Login
	err := c.ShouldBindXML(&xml)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if xml.User != "manu" || xml.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

// 绑定HTML表单的示例 (user=manu&password=123)
func BindHtml(c *gin.Context) {
	var form Login

	err := c.ShouldBind(&form)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if form.User != "manu" || form.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})

}
