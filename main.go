package main

import (
	"log"

	"github.com/aliasliao/naive-go-koa/addons/route"
	"github.com/aliasliao/naive-go-koa/core"
	"github.com/aliasliao/naive-go-koa/model"
)

func main() {
	router := route.NewRouter()
	router.Get("/user/:userId", func(ctx *core.Ctx) {
		userId := route.GetParam(ctx, "userId")
		ctx.SetCookie("sessionId", "80asd-dsd8-daf988das-88a0")
		ctx.Sendm(&model.User{
			Name:    userId,
			Age:     999,
			Hobbies: []string{"adsfa", "dddd"},
			Gender:  model.User_FEMALE,
		})
	}).Post("/user/:userId", func(ctx *core.Ctx) {
		userId := route.GetParam(ctx, "userId")
		user := &model.User{}
		ctx.Parsem(user)
		user.Name = userId
		ctx.Sendm(user)
	})

	app := core.New()
	app.Use(router)
	log.Println("app is listening on 3000...")
	err := app.Listen(3000)
	if err != nil {
		log.Fatal(err)
	}
}

// curl -v -H "Content-Type: application/json;charset=UTF-8" -d '{\"name\":\"123\",\"age\":999,\"hobbies\":[\"adsfa\",\"dddd\"],\"gender\":\"FEMALE\"}' localhost/user/123
