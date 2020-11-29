package core

import (
	"context"
	"net/http"
	"strconv"
)

type Handler func(*Ctx)

type Kao struct {
	server  *http.Server
	handler Handler
}

type Middleware interface {
	Apply(handler Handler) Handler
}

func New() *Kao {
	return &Kao{
		server: &http.Server{
			Addr: ":8080",
		},
	}
}

func (k *Kao) Use(m Middleware) *Kao {
	k.handler = m.Apply(k.handler)
	return k
}

func (k Kao) Listen(port int) error {
	k.server.Addr = ":" + strconv.Itoa(port)
	ctx := &Ctx{}
	k.server.Handler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		c, cancel := context.WithCancel(context.Background())
		ctx.Writer = writer
		ctx.Request = request
		ctx.Ctx = c

		k.handler(ctx)

		cancel()
	})
	return k.server.ListenAndServe()
}
