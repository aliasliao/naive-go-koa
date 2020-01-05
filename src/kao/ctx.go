package kao

import (
	"fmt"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

type Ctx struct {
	res   *http.ResponseWriter
	req   *http.Request
	param *map[string]string
	m     *jsonpb.Marshaler
}

func newCtx(res *http.ResponseWriter, req *http.Request, param *map[string]string) *Ctx {
	return &Ctx{
		res:   res,
		req:   req,
		param: param,
		m:     &jsonpb.Marshaler{},
	}
}

func (ctx Ctx) GetParam(key string) (string, bool) {
	v, ok := (*ctx.param)[key]
	return v, ok
}

func (ctx Ctx) GetQuery(key string) ([]string, bool) {
	v, ok := ctx.req.URL.Query()[key]
	return v, ok
}

func (ctx Ctx) Send(data ...interface{}) (n int, err error) {
	return fmt.Fprint(*ctx.res, data...)
}

func (ctx Ctx) Sendf(format string, data ...interface{}) (n int, err error) {
	return fmt.Fprintf(*ctx.res, format, data...)
}

func (ctx Ctx) SendMessage(pb proto.Message) error {
	ctx.SetHeader("Content-Type", "application/json;charset=UTF-8")
	return ctx.m.Marshal(*ctx.res, pb)
}

func (ctx Ctx) SetHeader(key string, val string) {
	(*ctx.res).Header().Set(key, val)
}

func (ctx Ctx) SetHeaders(headers map[string]string) {
	for key := range headers {
		ctx.SetHeader(key, headers[key])
	}
}

func (ctx Ctx) SetCookie(name string, value string) {
	http.SetCookie(*ctx.res, &http.Cookie{
		Name:  name,
		Value: value,
	})
}

func (ctx Ctx) SetCookies(cookies map[string]string) {
	for name := range cookies {
		ctx.SetCookie(name, cookies[name])
	}
}
