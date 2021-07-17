package controller

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

type ProfileForm struct {
	Name   string                `form:"name" binding:"required"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
	// 或多个文件
	// Avatars []*multipart.FileHeader `form:"avatar" binding:"required"`
}

func BindMultipart(c *gin.Context) {
	// 你可以使用显示绑定声明来绑定多部分表单:
	// c.ShouldBindWith(&form, binding.Form)
	// 或者你可以简单地将自动绑定与ShouldBind方法一起使用:
	var form ProfileForm
	// 在这种情况下，将自动选择适当的绑定
	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	err := c.SaveUploadedFile(form.Avatar, form.Avatar.Filename)
	if err != nil {
		c.String(http.StatusInternalServerError, "unknown error")
		return
	}
	// db.Save(&form)
	c.String(http.StatusOK, "ok")

}
