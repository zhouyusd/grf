package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Response struct {
	status int
	ctx    *gin.Context

	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Path      string      `json:"path,omitempty"`
	Timestamp int64       `json:"timestamp"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		status:    http.StatusOK,
		ctx:       ctx,
		Code:      int(StatusSuccess),
		Data:      nil,
		Msg:       "",
		Timestamp: time.Now().Unix(),
	}
}

func (hr *Response) SetStatus(status int) *Response {
	hr.status = status
	return hr
}

func (hr *Response) SetCode(code StatusCode) *Response {
	hr.Code = int(code)
	return hr
}

func (hr *Response) SetMessage(msg string) *Response {
	hr.Msg = msg
	return hr
}

func (hr *Response) SetPath(path string) *Response {
	hr.Path = path
	return hr
}

func (hr *Response) SetData(data interface{}) *Response {
	hr.Data = data
	return hr
}

func (hr *Response) JsonResponse() {
	hr.ctx.JSON(hr.status, hr)
}
