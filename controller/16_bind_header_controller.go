package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type testHeader struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
}

func BindHeader(c *gin.Context) {
	h := testHeader{}
	if err := c.ShouldBindHeader(&h); err != nil {
		c.JSON(http.StatusOK, err)
	}
	fmt.Printf("%#v\n", h)
	c.JSON(http.StatusOK, gin.H{"Rate": h.Rate, "Domain": h.Domain})
}
