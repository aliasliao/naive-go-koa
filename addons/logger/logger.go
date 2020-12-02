package logger

import (
	"log"

	"github.com/aliasliao/naive-go-koa/core"
)

func NewLogger() core.Middleware {
	return core.MiddlewareFunc(func(handler core.Handler) core.Handler {
		return func(ctx *core.Ctx) {
			method, path := ctx.Request.Method, ctx.Request.URL.Path
			log.Printf("-> %s %s\n", method, path)
			if handler != nil {
				handler(ctx)
			}
			log.Printf("<- %s %s %d\n", method, path, ctx.StatusCode)
		}
	})
}
