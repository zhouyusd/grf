package grf

import "github.com/gin-gonic/gin"

type HandlerFunc func(*Context)

type Engine struct {
	engine *gin.Engine
}

func handlerToGinHandler(handlers []HandlerFunc) []gin.HandlerFunc {
	ginHandlers := make([]gin.HandlerFunc, 0, len(handlers))
	for _, handle := range handlers {
		ginHandlers = append(ginHandlers, func(ginCtx *gin.Context) {
			handle(fromGinContext(ginCtx))
		})
	}
	return ginHandlers
}

func (e *Engine) GET(relativePath string, handlers ...HandlerFunc) *Engine {
	ginHandlers := handlerToGinHandler(handlers)
	e.engine.GET(relativePath, ginHandlers...)
	return e
}

func (e *Engine) POST(relativePath string, handlers ...HandlerFunc) *Engine {
	ginHandlers := handlerToGinHandler(handlers)
	e.engine.POST(relativePath, ginHandlers...)
	return e
}

func (e *Engine) PUT(relativePath string, handlers ...HandlerFunc) *Engine {
	ginHandlers := handlerToGinHandler(handlers)
	e.engine.PUT(relativePath, ginHandlers...)
	return e
}

func (e *Engine) PATCH(relativePath string, handlers ...HandlerFunc) *Engine {
	ginHandlers := handlerToGinHandler(handlers)
	e.engine.PATCH(relativePath, ginHandlers...)
	return e
}

func (e *Engine) DELETE(relativePath string, handlers ...HandlerFunc) *Engine {
	ginHandlers := handlerToGinHandler(handlers)
	e.engine.DELETE(relativePath, ginHandlers...)
	return e
}

func (e *Engine) OPTIONS(relativePath string, handlers ...HandlerFunc) *Engine {
	ginHandlers := handlerToGinHandler(handlers)
	e.engine.OPTIONS(relativePath, ginHandlers...)
	return e
}

func (e *Engine) HEAD(relativePath string, handlers ...HandlerFunc) *Engine {
	ginHandlers := handlerToGinHandler(handlers)
	e.engine.HEAD(relativePath, ginHandlers...)
	return e
}

func (e *Engine) Use(handlers ...HandlerFunc) {
	ginHandlers := handlerToGinHandler(handlers)
	e.engine.Use(ginHandlers...)
}

func (e *Engine) AddViews(views ...View) *Engine {
	for _, view := range views {
		view.Router(e)
	}
	return e
}

func (e *Engine) Run(addr ...string) error {
	return e.engine.Run(addr...)
}

func Default() *Engine {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	_ = engine.SetTrustedProxies(nil)
	return &Engine{engine: engine}
}
