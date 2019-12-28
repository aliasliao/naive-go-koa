package kao

import (
	"io"
	"log"
	"net/http"
	"strconv"
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
		log.Printf(
			"%s %s\n",
			req.Method,
			req.URL.Path,
		)
		method, url := req.Method, req.URL
		routeHit := false
		for _, route := range router.routes {
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
	})
	return &k
}

func (k Kao) Listen(port int, cb func(*string)) error {
	k.server.Addr = ":" + strconv.Itoa(port)
	cb(&(k.server.Addr))
	return k.server.ListenAndServe()
}
