package grf

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouyusd/grf/request"
	"github.com/zhouyusd/grf/response"
	"sync"
)

var (
	authorizer request.Authorizer = nil
	lock       sync.RWMutex
)

func SetAuthenticator(Authorizer request.Authorizer) {
	lock.Lock()
	defer lock.Unlock()
	authorizer = Authorizer
}

type Context struct {
	*gin.Context
	Request  *request.Request
	Response *response.Response
}

func fromGinContext(ctx *gin.Context) *Context {
	lock.RLock()
	authorize := authorizer
	lock.RUnlock()
	return &Context{
		Context:  ctx,
		Request:  request.NewRequest(ctx, authorize),
		Response: response.NewResponse(ctx),
	}
}
