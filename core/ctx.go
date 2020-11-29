package core

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

var (
	marshaler   = &jsonpb.Marshaler{EmitDefaults: true}
	unmarshaler = &jsonpb.Unmarshaler{}
)

type Ctx struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Ctx     context.Context
}

func (ctx *Ctx) GetQuery(key string) ([]string, bool) {
	v, ok := ctx.Request.URL.Query()[key]
	return v, ok
}

func (ctx *Ctx) Send(data ...interface{}) (n int, err error) {
	ctx.SetHeader("Content-Type", "text/plain;charset=UTF-8")
	return fmt.Fprint(ctx.Writer, data...)
}

func (ctx *Ctx) Sendf(format string, data ...interface{}) (n int, err error) {
	ctx.SetHeader("Content-Type", "text/plain;charset=UTF-8")
	return fmt.Fprintf(ctx.Writer, format, data...)
}

func (ctx *Ctx) Sendm(pb proto.Message) error {
	ctx.SetHeader("Content-Type", "application/json;charset=UTF-8")
	return marshaler.Marshal(ctx.Writer, pb)
}

func (ctx *Ctx) Parsem(pb proto.Message) error {
	return unmarshaler.Unmarshal(ctx.Request.Body, pb)
}

func (ctx *Ctx) SetHeader(key string, val string) {
	ctx.Writer.Header().Set(key, val)
}

func (ctx *Ctx) SetHeaders(headers map[string]string) {
	for key := range headers {
		ctx.SetHeader(key, headers[key])
	}
}

func (ctx *Ctx) SetCookie(name string, value string) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:  name,
		Value: value,
	})
}

func (ctx *Ctx) SetCookies(cookies map[string]string) {
	for name := range cookies {
		ctx.SetCookie(name, cookies[name])
	}
}
