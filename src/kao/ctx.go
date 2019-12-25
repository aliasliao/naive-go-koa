package kao

import (
	"fmt"
	"net/http"
)

type Ctx struct {
	res   *http.ResponseWriter
	req   *http.Request
	param *map[string]string
}

func newCtx(res *http.ResponseWriter, req *http.Request, param *map[string]string) *Ctx {
	return &Ctx{
		res:   res,
		req:   req,
		param: param,
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

func (ctx Ctx) Send(data interface{}) {
	_, _ = fmt.Fprint(*ctx.res, data)
}
