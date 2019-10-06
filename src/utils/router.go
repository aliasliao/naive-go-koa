package router

import (
	"io"
	"log"
	"net/http"
)

type Action = func(http.ResponseWriter, *http.Request)
type Mux = map[string]map[string]Action

type Route struct {
	Routes Mux
}

func New() *Route {
	mux := make(Mux, 0)
	mux["*"]["*"] = func(w http.ResponseWriter, r *http.Request) {
		if _, e := io.WriteString(w, "hello go-server"); e != nil {
			log.Fatal(e)
		}
	}
	return &Route{mux}
}

func (route *Route) Get(pattern string, handle Action) *Route {
	route.Routes["GET"][pattern] = handle
	return route
}

func (route *Route) Post(pattern string, handle Action) *Route {
	route.Routes["POST"][pattern] = handle
	return route
}

func (route *Route) Delete(pattern string, handle Action) *Route {
	route.Routes["DELETE"][pattern] = handle
	return route
}

func (route *Route) Put(pattern string, handle Action) *Route {
	route.Routes["PUT"][pattern] = handle
	return route
}

func (route *Route) Patch(pattern string, handle Action) *Route {
	route.Routes["PATCH"][pattern] = handle
	return route
}
