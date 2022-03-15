package grf

func RateThrottle() HandlerFunc {
	return func(ctx *Context) {
		ctx.Ctx.Next()
	}
}
