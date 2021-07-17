package controller

import "github.com/gin-gonic/gin"

// https://github.com/gin-gonic/gin/issues/129#issuecomment-124260092

type myForm struct {
	Colors []string `form:"colors[]"`
}

func IndexHandler(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}

func FormHandler(c *gin.Context) {
	var fakeForm myForm
	err := c.Bind(&fakeForm)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}
