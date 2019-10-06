package router

import (
	"io"
	"log"
	"net/http"
)

type Action = func(http.ResponseWriter, *http.Request)
type Dict = map[string]Action
type Mux = map[Method]Dict

type Route struct {
	Routes Mux
}

func New() *Route {
	mux := make(Mux, 0)
	mux[ANY] = Dict{
		"*": func(w http.ResponseWriter, r *http.Request) {
			if _, e := io.WriteString(w, "hello go-server"); e != nil {
				log.Fatal(e)
			}
		}}
	for _, method := range []Method{GET, POST, DELETE, PUT, PATCH} {
		mux[method] = make(Dict, 0)
	}
	return &Route{mux}
}

func (route *Route) Get(pattern string, handle Action) *Route {
	route.Routes[GET][pattern] = handle
	return route
}

func (route *Route) Post(pattern string, handle Action) *Route {
	route.Routes[POST][pattern] = handle
	return route
}

func (route *Route) Delete(pattern string, handle Action) *Route {
	route.Routes[DELETE][pattern] = handle
	return route
}

func (route *Route) Put(pattern string, handle Action) *Route {
	route.Routes[PUT][pattern] = handle
	return route
}

func (route *Route) Patch(pattern string, handle Action) *Route {
	route.Routes[PATCH][pattern] = handle
	return route
}
