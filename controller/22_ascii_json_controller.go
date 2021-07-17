package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//	使用 AsciiJSON 生成仅有 ASCII 字符的 JSON，非 ASCII 字符将会被转义 。
func AsciiJSON(c *gin.Context) {
	data := gin.H{
		"lang": "GO语言",
		"tag":  "<br>",
	}
	// 将会输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	c.AsciiJSON(http.StatusOK, data)
}
