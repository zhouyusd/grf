package grf

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type StatusCode int

const (
	StatusSuccess StatusCode = 20000

	StatusParameterInvalid    StatusCode = 40000
	StatusNotLogin            StatusCode = 40100
	StatusFrequently          StatusCode = 40200
	StatusForbidden           StatusCode = 40300
	StatusNotFound            StatusCode = 40400
	StatusTokenInvalid        StatusCode = 40500
	StatusOtherClientLoggedIn StatusCode = 40600
	StatusSystemInternal      StatusCode = 50000
	StatusSystemTimeout       StatusCode = 50100
	StatusEtcdConnFailed      StatusCode = 52379
	StatusMysqlConnFailed     StatusCode = 53306
	StatusRedisConnFailed     StatusCode = 56379
)

type Response struct {
	status int
	ctx    *gin.Context

	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Path      string      `json:"path"`
	Timestamp int64       `json:"timestamp"`
}

func makeResponse(ctx *gin.Context) *Response {
	return &Response{
		status:    http.StatusOK,
		ctx:       ctx,
		Code:      int(StatusSuccess),
		Data:      nil,
		Msg:       "",
		Path:      ctx.Request.RequestURI,
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

func (hr *Response) SetData(data interface{}) *Response {
	hr.Data = data
	return hr
}

func (hr *Response) JsonResponse() {
	hr.ctx.JSON(hr.status, hr)
}
