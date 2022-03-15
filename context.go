package grf

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Context struct {
	Ctx      *gin.Context
	Request  *Request
	Response *Response
}

func fromGinContext(ctx *gin.Context) *Context {
	return &Context{
		Ctx:      ctx,
		Request:  makeRequest(ctx),
		Response: makeResponse(ctx),
	}
}

/************************************/
/************* CONTEXT **************/
/************************************/

// Deadline always returns that there is no deadline (ok==false),
// maybe you want to use Request.Context().Deadline() instead.
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return
}

// Done always returns nil (chan which will wait forever),
// if you want to abort your work when the connection was closed
// you should use Request.Context().Done() instead.
func (c *Context) Done() <-chan struct{} {
	return nil
}

// Err always returns nil, maybe you want to use Request.Context().Err() instead.
func (c *Context) Err() error {
	return nil
}

// Value returns the value associated with this context for key, or nil
// if no value is associated with key. Successive calls to Value with
// the same key returns the same result.
func (c *Context) Value(key interface{}) interface{} {
	return c.Ctx.Value(key)
}
