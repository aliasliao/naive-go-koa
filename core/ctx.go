package core

import (
	"context"
	"fmt"
	"net/http"
)

type Ctx struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	StatusCode int
	Ctx        context.Context
}

func (ctx *Ctx) SetHeader(key string, val string) {
	ctx.Writer.Header().Set(key, val)
}

func (ctx *Ctx) GetQuery(key string) ([]string, bool) {
	v, ok := ctx.Request.URL.Query()[key]
	return v, ok
}

func (ctx *Ctx) SetCookie(name string, value string) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:  name,
		Value: value,
	})
}

func (ctx *Ctx) Error(detail string, statusCode int) {
	http.Error(ctx.Writer, detail, statusCode)
	ctx.StatusCode = statusCode
}

func (ctx *Ctx) Message(message string, contentType string) {
	if contentType == "" {
		contentType = "application/json;charset=UTF-8"
	}
	ctx.SetHeader("Content-Type", contentType)
	if _, err := fmt.Fprintln(ctx.Writer, message); err != nil {
		ctx.Error(err.Error(), http.StatusInternalServerError)
		ctx.StatusCode = http.StatusInternalServerError
	}
	ctx.StatusCode = http.StatusOK
}

func (ctx *Ctx) Write(message []byte, contentType string) {
	if contentType == "" {
		contentType = "text/plain;charset=UTF-8"
	}
	ctx.SetHeader("Content-Type", contentType)
	if _, err := ctx.Writer.Write(message); err != nil {
		ctx.Error(err.Error(), http.StatusInternalServerError)
		ctx.StatusCode = http.StatusInternalServerError
	}
	ctx.StatusCode = http.StatusOK
}
