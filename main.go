package main

import (
	"github.com/gin-gonic/gin"
	_ "go-gin/config"
	"go-gin/db"
	"go-gin/routers"
)

/**
应用程序启动入口
*/
func main() {

	// 第一步：设置数据库
	err := db.InitMySQL()
	if err != nil {
		panic(err)
	}
	// 程序退出关闭数据库连接
	defer db.Close()

	//	第二步：设置HTTP路由
	//	初始化路由
	//	包括各种http的功能API
	engine := gin.Default()
	routers.InitRouter(engine)

	//	第三步：设置web服务端口
	_ = engine.Run(":9999")
}
