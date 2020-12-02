package router

import (
	"context"
	"log"
	"regexp"

	"github.com/aliasliao/naive-go-koa/addons/router/pathToRegexp"
	"github.com/aliasliao/naive-go-koa/core"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
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

var (
	marshaler   = &jsonpb.Marshaler{EmitDefaults: true}
	unmarshaler = &jsonpb.Unmarshaler{}
)

type routeT struct {
	re      *regexp.Regexp
	method  Method
	handler core.Handler
}

// Router is a controller middleware
//
// Example:
//  r := router.NewRouter()
//  r.Get("/user/:userId", func(ctx *core.Ctx) {
//  	userId := router.GetParam(ctx, "userId")
//  	ctx.SetCookie("sessionId", "80asd-dsd8-daf988das-88a0")
//  	router.Sendm(ctx, &model.User{
//  		Name:    userId,
//  		Age:     999,
//  		Hobbies: []string{"adsfa", "dddd"},
//  		Gender:  model.User_FEMALE,
//  	})
//  }).Post("/user/:userId", func(ctx *core.Ctx) {
//  	userId := router.GetParam(ctx, "userId")
//  	user := &model.User{}
//  	router.Parsem(ctx, user)
//  	user.Name = userId
//  	router.Sendm(ctx, user)
//  })
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

func Sendm(ctx *core.Ctx, pb proto.Message) error {
	ctx.SetHeader("Content-Type", "application/json;charset=UTF-8")
	return marshaler.Marshal(ctx.Writer, pb)
}

func Parsem(ctx *core.Ctx, pb proto.Message) error {
	return unmarshaler.Unmarshal(ctx.Request.Body, pb)
}

// implement core.Handler
func (r *Router) Apply(handler core.Handler) core.Handler {
	return func(ctx *core.Ctx) {
		method, path := ctx.Request.Method, ctx.Request.URL.Path
		log.Printf("%s %s\n", method, path)
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
				break
			}
		}
		if handler != nil {
			handler(ctx)
		}
	}
}
