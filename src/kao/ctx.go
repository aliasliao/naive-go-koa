package kao

import (
	"net/http"
	"regexp"
)

type Ctx struct {
	res *http.ResponseWriter
	req *http.Request
	re  *regexp.Regexp
}

func newCtx(res *http.ResponseWriter, req *http.Request, re *regexp.Regexp) *Ctx {
	return &Ctx{
		res: res,
		req: req,
		re:  re,
	}
}

func (ctx Ctx) GetParam(key string) string {}

func (ctx Ctx) GetQuery(key string) string {}

func (ctx Ctx) Send(data interface{}) {}
