package kao

import (
	"net/http"
	"strconv"
)

type Kao struct {
	server *http.Server
}

type middlewareT interface {
	Middleware(handler http.Handler) http.Handler
}

func New() *Kao {
	return &Kao{
		server: &http.Server{
			Addr:    ":8080",
			Handler: nil,
		},
	}
}

func (k Kao) Use(m middlewareT) *Kao {
	k.server.Handler = m.Middleware(k.server.Handler)
	return &k
}

func (k Kao) Listen(port int, cb func(*string)) error {
	k.server.Addr = ":" + strconv.Itoa(port)
	cb(&(k.server.Addr))
	return k.server.ListenAndServe()
}
