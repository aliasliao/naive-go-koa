package main

import (
	"flag"
	"log"

	"github.com/aliasliao/naive-go-koa/addons/router"
	"github.com/aliasliao/naive-go-koa/addons/serve"
	"github.com/aliasliao/naive-go-koa/core"
	"github.com/aliasliao/naive-go-koa/model"
)

var (
	port = flag.Int("port", 3000, "Server port")
	dir  = flag.String("dir", ".", "Public path")
)

func main() {
	flag.Parse()

	r := router.NewRouter()
	r.Get("/user/:userId", func(ctx *core.Ctx) {
		userId := router.GetParam(ctx, "userId")
		ctx.SetCookie("sessionId", "80asd-dsd8-daf988das-88a0")
		ctx.Sendm(&model.User{
			Name:    userId,
			Age:     999,
			Hobbies: []string{"adsfa", "dddd"},
			Gender:  model.User_FEMALE,
		})
	}).Post("/user/:userId", func(ctx *core.Ctx) {
		userId := router.GetParam(ctx, "userId")
		user := &model.User{}
		ctx.Parsem(user)
		user.Name = userId
		ctx.Sendm(user)
	})

	app := core.New()
	app.Use(r)
	if s, err := serve.New(*dir); err != nil {
		log.Fatal(err)
	} else {
		app.Use(s)
	}
	log.Printf("app is listening on %d...", *port)
	if err := app.Listen(*port); err != nil {
		log.Fatal(err)
	}
}

// curl -v -H "Content-Type: application/json;charset=UTF-8" -d '{\"name\":\"123\",\"age\":999,\"hobbies\":[\"adsfa\",\"dddd\"],\"gender\":\"FEMALE\"}' localhost/user/123
