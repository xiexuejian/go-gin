package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SecureJson 你也可以使用自己的secure json前缀
// r.SecureJsonPrefix(")]}',\n")
func SecureJson(c *gin.Context) {

	names := []string{"lena", "austin", "foo"}

	// 将会输出 : while(1);["lena","austin","foo"]
	c.SecureJSON(http.StatusOK, names)
}
