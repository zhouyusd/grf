package grf

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type Request struct {
	ctx   *gin.Context
	User  interface{}
	Token string
}

type Authorizer func(*gin.Context) (interface{}, string)

var (
	authorizer Authorizer = nil
	lock       sync.RWMutex
)

func SetAuthenticator(Authorizer Authorizer) {
	lock.Lock()
	defer lock.Unlock()
	authorizer = Authorizer
}

func makeRequest(ctx *gin.Context) *Request {
	var user interface{}
	var token string
	lock.RLock()
	if authorizer != nil {
		user, token = authorizer(ctx)
	}
	lock.RUnlock()
	return &Request{
		ctx:   ctx,
		User:  user,
		Token: token,
	}
}
