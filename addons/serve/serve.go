package serve

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aliasliao/naive-go-koa/core"
)

func New(dir string) (core.Middleware, error) {
	fileInfo, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}
	if !fileInfo.IsDir() {
		return nil, errors.New(dir + " is not a valid directory")
	}
	return core.MiddlewareFunc(func(handler core.Handler) core.Handler {
		return func(ctx *core.Ctx) {
			method, path := ctx.Request.Method, ctx.Request.URL.Path
			if method == http.MethodGet {
				file := filepath.Join(dir, path)
				if _, err := os.Stat(file); err != nil {
					ctx.Error(err.Error(), http.StatusNotFound)
				}
				if data, err := ioutil.ReadFile(file); err != nil {
					ctx.Error(err.Error(), http.StatusBadRequest)
				} else {
					ctx.Write(data)
				}
			}
			if handler != nil {
				handler(ctx)
			}
		}
	}), nil
}
