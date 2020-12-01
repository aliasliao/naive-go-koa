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

func (ctx *Ctx) setHeader(key string, val string) {
	ctx.Writer.Header().Set(key, val)
}

func (ctx *Ctx) GetQuery(key string) ([]string, bool) {
	v, ok := ctx.Request.URL.Query()[key]
	return v, ok
}

func (ctx *Ctx) Sendm(pb proto.Message) error {
	ctx.setHeader("Content-Type", "application/json;charset=UTF-8")
	return marshaler.Marshal(ctx.Writer, pb)
}

func (ctx *Ctx) Parsem(pb proto.Message) error {
	return unmarshaler.Unmarshal(ctx.Request.Body, pb)
}

func (ctx *Ctx) SetCookie(name string, value string) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:  name,
		Value: value,
	})
}

func (ctx *Ctx) Error(detail string, statusCode int) {
	http.Error(ctx.Writer, detail, statusCode)
}

func (ctx *Ctx) Message(message string) {
	ctx.setHeader("Content-Type", "application/json;charset=UTF-8")
	if _, err := fmt.Fprintln(ctx.Writer, message); err != nil {
		ctx.Error(err.Error(), http.StatusInternalServerError)
	}
}

func (ctx *Ctx) Write(message []byte) {
	ctx.setHeader("Content-Type", "text/plain;charset=UTF-8")
	if _, err := ctx.Writer.Write(message); err != nil {
		ctx.Error(err.Error(), http.StatusInternalServerError)
	}
}
