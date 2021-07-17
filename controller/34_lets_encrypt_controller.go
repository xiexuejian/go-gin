package controller

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

//	第一个案例
//import (
//	"github.com/gin-gonic/autotls"
//	"github.com/gin-gonic/gin"
//	"log"
//)
//func test() {
//	r := gin.Default()
//	// Ping handler
//	r.GET("/ping", func(c *gin.Context) {
//		c.String(200, "pong")
//	})
//	log.Fatal(autotls.Run(r, "example1.com", "example2.com"))
//}

//	第二个案例
//	自定义 autocert 管理器示例。
import (
	"log"
)

func main() {
	r := gin.Default()
	// Ping 处理器
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}
	log.Fatal(autotls.RunWithManager(r, &m))
}
