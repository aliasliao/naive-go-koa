package kao

import (
	"regexp"

	"pathToRegexp"
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

func registerRoute(r Router, method Method, path string, ctr func(*Ctx)) *Router {
	re := pathToRegexp.PathToRegexp(path, nil)
	route := &routeT{
		re:      re,
		method:  &method,
		handler: &ctr,
	}
	r.routes = append(r.routes, route)
	return &r
}
func (r Router) Get(path string, ctr func(*Ctx)) *Router {
	return registerRoute(r, GET, path, ctr)
}
func (r Router) Post(path string, ctr func(*Ctx)) *Router {
	return registerRoute(r, POST, path, ctr)
}
func (r Router) Put(path string, ctr func(*Ctx)) *Router {
	return registerRoute(r, PUT, path, ctr)
}
func (r Router) Patch(path string, ctr func(*Ctx)) *Router {
	return registerRoute(r, PATCH, path, ctr)
}
func (r Router) Delete(path string, ctr func(*Ctx)) *Router {
	return registerRoute(r, DELETE, path, ctr)
}
