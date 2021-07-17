package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//	上传文件的文件名可以由用户自定义，所以可能包含非法字符串，为了安全起见，应该由服务端统一文件名规则。

/**
 *	参考：
 *     https://github.com/gin-gonic/gin/issues/774
 * 	   https://github.com/gin-gonic/gin/issues/1693
 *	   https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition#directives
 */
func SimpleFile(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")
	log.Infoln(file.Filename)
	// 上传文件到指定的路径
	err := c.SaveUploadedFile(file, "/file")
	if err != nil {
		errorInfo := fmt.Errorf("上传失败")
		//	自定义状态码，-3001
		c.String(-3001, errorInfo.Error())
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func MultiFile(c *gin.Context) {
	// 多文件
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	for _, file := range files {
		log.Println(file.Filename)
		//上传文件到指定的路径
		err := c.SaveUploadedFile(file, "/file")
		if err != nil {
			errorInfo := fmt.Errorf("上传失败")
			//	自定义状态码，-3001
			c.String(-3001, errorInfo.Error())
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
