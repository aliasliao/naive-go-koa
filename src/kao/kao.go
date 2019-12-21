package kao

import (
	"io"
	"net/http"
)

type Kao struct {
	server *http.Server
}

func New() *Kao {
	return &Kao{
		server: &http.Server{
			Addr:    ":8080",
			Handler: nil,
		},
	}
}

func (k Kao) Use(router *Router) *Kao {
	k.server.Handler = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		method, url := req.Method, req.URL
		routeHit := false
		for _, route := range router.routes {
			if *route.method == Method(method) && route.re.MatchString(url.Path) {
				(*route.handler)(newCtx(&res, req, route.re))
				routeHit = true
				break
			}
		}
		if !routeHit {
			io.WriteString(res, "Using Default Response Handler~")
		}
	})
	return &k
}

func (k Kao) Listen(port string) {

}
