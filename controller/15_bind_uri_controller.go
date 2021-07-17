package controller

import "github.com/gin-gonic/gin"

//	https://github.com/gin-gonic/gin/issues/846

type Person3 struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func BindUri(c *gin.Context) {
	var person Person3

	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
}
