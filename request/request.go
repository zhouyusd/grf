package request

import (
	"github.com/gin-gonic/gin"
)

type Request struct {
	ctx   *gin.Context
	User  interface{}
	Token string
}

type Authorizer func(*gin.Context) (interface{}, string)

func NewRequest(ctx *gin.Context, authorize Authorizer) *Request {
	var user interface{}
	var token string
	if authorize != nil {
		user, token = authorize(ctx)
	}
	return &Request{
		ctx:   ctx,
		User:  user,
		Token: token,
	}
}
