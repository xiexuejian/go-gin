package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"unsafe"
)

type Data struct {
	Id      string
	Page    string
	Name    string
	Message string
}

//	http://localhost:9999/getAndPost?id=1&page=10
func GetAndPost(c *gin.Context) {
	//	get
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	//	post
	name := c.PostForm("name")
	message := c.PostForm("message")

	data := &Data{
		Id:      id,
		Page:    page,
		Name:    name,
		Message: message,
	}
	result, _ := json.Marshal(data)
	c.String(http.StatusOK, *(*string)(unsafe.Pointer(&result)))
	//	{"Id":"1","Page":"10","Name":"2222","Message":"1111"}
}
