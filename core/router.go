package core

import (
	"io"
	"log"
	"net/http"
	"regexp"

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

type routeT struct {
	re      *regexp.Regexp
	method  *Method
	handler *func(*Ctx)
}
type Router struct {
	routes []*routeT
}

func NewRouter() *Router {
	return &Router{
		routes: make([]*routeT, 0),
	}
}

func registerRoute(r *Router, method Method, path string, ctr func(*Ctx)) *Router {
	re := pathToRegexp.PathToRegexp(path, nil)
	route := &routeT{
		re:      re,
		method:  &method,
		handler: &ctr,
	}
	r.routes = append(r.routes, route)
	return r
}
func (r *Router) Get(path string, ctr func(*Ctx)) *Router {
	return registerRoute(r, GET, path, ctr)
}
func (r *Router) Post(path string, ctr func(*Ctx)) *Router {
	return registerRoute(r, POST, path, ctr)
}
func (r *Router) Put(path string, ctr func(*Ctx)) *Router {
	return registerRoute(r, PUT, path, ctr)
}
func (r *Router) Patch(path string, ctr func(*Ctx)) *Router {
	return registerRoute(r, PATCH, path, ctr)
}
func (r *Router) Delete(path string, ctr func(*Ctx)) *Router {
	return registerRoute(r, DELETE, path, ctr)
}

func (r Router) Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Printf(
			"%s %s\n",
			req.Method,
			req.URL.Path,
		)
		method, url := req.Method, req.URL
		routeHit := false
		for _, route := range r.routes {
			if *route.method == Method(method) && route.re.MatchString(url.Path) {
				keys := route.re.SubexpNames()
				values := route.re.FindAllStringSubmatch(url.Path, -1)[0]
				param := map[string]string{}
				for i, key := range keys {
					param[key] = values[i]
				}
				(*route.handler)(newCtx(&res, req, &param))
				routeHit = true
				break
			}
		}
		if !routeHit {
			_, _ = io.WriteString(res, "Using Default Response Handler~")
		}
		if handler != nil {
			handler.ServeHTTP(res, req)
		}
	})
}
