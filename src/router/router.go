package router

import (
	"io"
	"log"
	"net/http"
)

type Action = func(http.ResponseWriter, *http.Request)
type Dict = map[string]interface{}
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

func (route *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method, url := StringToMethod(r.Method), r.URL.String()
	fallbackAction := route.Routes[ANY]["*"].(Action)
	if dict, exist1 := route.Routes[method]; exist1 {
		if action, exist2 := dict[url]; exist2 {
			if url == "/api/test" && method == POST {
				// interception here
				// Reference: https://coderwall.com/p/b5dpwq/fun-with-the-reflection-package-to-analyse-any-function
			} else {
				action.(Action)(w, r)
			}
		} else {
			fallbackAction(w, r)
		}
	} else {
		fallbackAction(w, r)
	}
}

func (route *Route) Get(pattern string, action interface{}) *Route {
	route.Routes[GET][pattern] = action
	return route
}

func (route *Route) Post(pattern string, action interface{}) *Route {
	route.Routes[POST][pattern] = action
	return route
}

func (route *Route) Delete(pattern string, action interface{}) *Route {
	route.Routes[DELETE][pattern] = action
	return route
}

func (route *Route) Put(pattern string, action interface{}) *Route {
	route.Routes[PUT][pattern] = action
	return route
}

func (route *Route) Patch(pattern string, action interface{}) *Route {
	route.Routes[PATCH][pattern] = action
	return route
}
