package grf

import (
	"github.com/gin-gonic/gin"
)

func defaultLogger(c *Context) {}

func defaultRecovery(ctx *Context) {
	recoveryHandler := gin.CustomRecoveryWithWriter(nil, func(*gin.Context, interface{}) {
		ctx.Response.
			SetCode(StatusSystemInternal).
			SetMessage("系统错误,请稍后再试～").
			SetData(nil).JsonResponse()
	})
	recoveryHandler(ctx.Ctx)
}
