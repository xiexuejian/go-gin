package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//	绑定 request body 的常规方法是使用 c.Request.Body 并且不能多次调用它们。

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}
type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func ShouldBindDemo(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// This c.ShouldBind consumes c.Request.Body and it cannot be reused.
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// Always an error is occurred by this because c.Request.Body is EOF now.
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		log.Infoln("其它种情况")
	}
}

//	还可以使用 c.ShouldBindBodyWith 。
func ShouldBindBodyWithDemo(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// 这里读取 c.Request.Body 并将结果存储到 context 中。
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// At this time, it reuses body stored in the context.
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, `the body should be formB JSON`)
		// And it can accepts other formats
	} else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
		c.String(http.StatusOK, `the body should be formB XML`)
	} else {
		log.Infoln("其它种情况")
	}
}

/**
c.ShouldBindBodyWith 在绑定前存储 body 到 context 中。这对性能会有轻微的影响，所以如
果你可以通过立即调用绑定，不应该使用这个方法。

只有一些格式需要这个功能 -- JSON 、 XML 、 MsgPack、 ProtoBuf 。 对于其他格式，
Query、 Form、 FormPost、 FormMultipart， 能被 c.ShouldBind() 多次调用，而不会对性能
造成任何损害 （参见 #1341）:https://github.com/gin-gonic/gin/pull/1341

*/
