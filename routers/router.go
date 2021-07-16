package routers

import (
	"github.com/gin-gonic/gin"
	. "go-gin/controller"
)

//	初始化路由
//	对应官方文档的API例子
func InitRouter(request *gin.Engine) {

	//	1、路由参数
	request.GET("/testGet/:name", TestGet)
	//	全匹配
	request.GET("/testGet/:name/*action", TestGet02)
	request.GET("/testGet03/:name/action", TestGet03)

	//	2、查询字符串参数
	request.GET("/testStringParam", TestGet04)

	//	3、Multipart/Urlencoded 表单
	request.POST("/formPost", FormPost)

	//	4、query（get）+post 表单
	request.POST("/getAndPost", GetAndPost)

	//	5、Map 作为查询字符串或 post表单 参数
	request.POST("/mapPost", MapAndPost)

	//	6、上传文件
	// 给表单限制上传大小 (默认是 32 MiB)
	request.MaxMultipartMemory = 8 << 20 // 8 MiB
	//	单个文件
	request.POST("/simpleFile", SimpleFile)
	//	多个文件
	request.POST("/multiFile", MultiFile)
}
