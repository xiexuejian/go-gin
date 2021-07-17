package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

//	在中间件或处理程序中启动新的Goroutines时，你不应该使用其中的原始上下文，你必须使用只读副本
//（c.Copy()）
func LongAsync(c *gin.Context) {

	// 创建在goroutine中使用的副本
	cCp := c.Copy()
	go func() {
		// 使用time.sleep()休眠5秒，模拟一个用时长的任务
		time.Sleep(5 * time.Second)
		// 注意，你使用的是复制的 context "cCp" ，重要
		log.Println("Done! in path " + cCp.Request.URL.Path)
	}()
}

func LongSync(c *gin.Context) {

	// 使用 time.Sleep() 休眠 5 秒，模拟一个用时长的任务。
	time.Sleep(5 * time.Second)
	// 因为我们没有使用协程，我们不需要复制 context
	log.Println("Done! in path " + c.Request.URL.Path)
}
