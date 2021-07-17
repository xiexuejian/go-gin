package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LocalFile(c *gin.Context) {
	c.File("local/file.go")
}

func FsFile(c *gin.Context) {
	var fs http.FileSystem // 需要指定，给其赋值
	c.FileFromFS("fs/file.go", fs)
}
