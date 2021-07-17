package gin_log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

/**
 * 关于gin的日志问题
 */

func init() {

	//  知识点一：禁用控制台颜色，当你将日志写入到文件的时候，你不需要控制台颜色
	gin.DisableConsoleColor()

	//	知识点二：写入日志文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 知识点三：如果你需要同时写入日志文件和控制台上显示，使用下面代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//	知识点四：自定义日志格式
	router := gin.New()
	// LoggerWithFormatter 中间件会将日志写入 gin.DefaultWriter
	// 默认 gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(define))
	router.Use(gin.Recovery())

	//	知识点五：控制日志输出颜色（Controlling Log output coloring）
	//记录日志的颜色
	gin.ForceConsoleColor()

}

func define(param gin.LogFormatterParams) string {
	// 你的自定义格式
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
