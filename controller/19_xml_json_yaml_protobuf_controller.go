package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func SomeJson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
}

func MoreJson(c *gin.Context) {
	// 你也可以使用一个 struct
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "Lena"
	msg.Message = "hey"
	msg.Number = 123
	// 注意 msg.Name 在json中会变成 "user"
	// 将会输出 : {"user": "Lena", "Message": "hey", "Number": 123}
	c.JSON(http.StatusOK, msg)

}

func SomeXml(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
}

func SomeYaml(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})

}

func SomeProtoBuf(c *gin.Context) {
	reps := []int64{int64(1), int64(2)}
	label := "test"
	// protobuf 的特定定义写在testdata/protoexample文件中
	data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
	}
	// 注意数据在响应中变为二进制数据
	// 将会输出protoexample.test protobuf序列化的数据
	c.ProtoBuf(http.StatusOK, data)
}
