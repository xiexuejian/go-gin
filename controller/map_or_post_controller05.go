package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MapAndPost(c *gin.Context) {
	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")
	fmt.Printf("ids: %v; names: %v\n", ids, names)
	c.String(http.StatusOK, "结果-->ids:%v		names:%v", ids, names)
	//	结果-->ids:map[a:1234 b:hello]		names:map[first:1111 second:2222]
}
