package koa

import (
	"regexp"

	"pathToRegexp"
)

type Method int

const (
	GET Method = iota
	POST
	PUT
	PATCH
	DELETE
	OPTION
)

type Ctx interface{}

type layerT struct {
	re      *regexp.Regexp
	methods *[]Method
	handler *func(Ctx)
}
type Router struct {
	layers []*layerT
}

func NewRouter() *Router {
	return &Router{
		layers: make([]*layerT, 0),
	}
}

func updateRouter(r Router, methods []Method, path string, ctr func(Ctx)) *Router {
	re := pathToRegexp.PathToRegexp(path, nil)
	layer := &layerT{
		re:      re,
		methods: &methods,
		handler: &ctr,
	}
	r.layers = append(r.layers, layer)
	return &r
}
func (r Router) Get(path string, ctr func(Ctx)) *Router {
	return updateRouter(r, []Method{GET}, path, ctr)
}
func (r Router) Post(path string, ctr func(Ctx)) *Router {
	return updateRouter(r, []Method{POST}, path, ctr)
}
func (r Router) Put(path string, ctr func(Ctx)) *Router {
	return updateRouter(r, []Method{PUT}, path, ctr)
}
func (r Router) Patch(path string, ctr func(Ctx)) *Router {
	return updateRouter(r, []Method{PATCH}, path, ctr)
}
func (r Router) Delete(path string, ctr func(Ctx)) *Router {
	return updateRouter(r, []Method{DELETE}, path, ctr)
}
func (r Router) Option(path string, ctr func(Ctx)) *Router {
	return updateRouter(r, []Method{OPTION}, path, ctr)
}
func (r Router) Register(methods []Method, path string, ctr func(Ctx)) *Router {
	return updateRouter(r, methods, path, ctr)
}
