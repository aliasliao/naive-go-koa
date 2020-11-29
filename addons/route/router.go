package route

import (
	"context"
	"log"
	"regexp"

	"github.com/aliasliao/naive-go-koa/core"
	"github.com/aliasliao/naive-go-koa/pathToRegexp"
)

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	PATCH  Method = "PATCH"
	DELETE Method = "DELETE"
)

type key string

const (
	paramsKey key = "params-key"
)

type routeT struct {
	re      *regexp.Regexp
	method  Method
	handler core.Handler
}
type Router struct {
	routes []*routeT
}

func NewRouter() *Router {
	return &Router{
		routes: make([]*routeT, 0),
	}
}

func (r *Router) registerRoute(method Method, path string, handler core.Handler) *Router {
	re := pathToRegexp.PathToRegexp(path, nil)
	route := &routeT{
		re:      re,
		method:  method,
		handler: handler,
	}
	r.routes = append(r.routes, route)
	return r
}
func (r *Router) Get(path string, handler core.Handler) *Router {
	return r.registerRoute(GET, path, handler)
}
func (r *Router) Post(path string, handler core.Handler) *Router {
	return r.registerRoute(POST, path, handler)
}
func (r *Router) Put(path string, handler core.Handler) *Router {
	return r.registerRoute(PUT, path, handler)
}
func (r *Router) Patch(path string, handler core.Handler) *Router {
	return r.registerRoute(PATCH, path, handler)
}
func (r *Router) Delete(path string, handler core.Handler) *Router {
	return r.registerRoute(DELETE, path, handler)
}

func GetParam(ctx *core.Ctx, key string) string {
	if params := ctx.Ctx.Value(paramsKey); params != nil {
		if p, ok := params.(*map[string]string); ok {
			return (*p)[key]
		}
	}
	return ""
}

// implement core.Handler
func (r Router) Apply(handler core.Handler) core.Handler {
	return func(ctx *core.Ctx) {
		method, path := ctx.Request.Method, ctx.Request.URL.Path
		log.Printf("%s %s\n", method, path)
		routeHit := false
		for _, route := range r.routes {
			if route.method == Method(method) && route.re.MatchString(path) {
				keys := route.re.SubexpNames()
				values := route.re.FindAllStringSubmatch(path, -1)[0]
				param := map[string]string{}
				for i, key := range keys {
					param[key] = values[i]
				}
				ctx.Ctx = context.WithValue(ctx.Ctx, paramsKey, &param)
				route.handler(ctx)
				routeHit = true
				break
			}
		}
		if !routeHit {
			// 404 Not Found
			ctx.Send("Using Default Response Handler~")
		}
		if handler != nil {
			// inner handler
			handler(ctx)
		}
	}
}
