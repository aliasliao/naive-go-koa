package main

import (
	"kao"
)

func main() {
	router := kao.NewRouter()
	router.Get("/user/:userId", func(ctx *kao.Ctx) {
		userId, _ := ctx.GetParam("userId")
		println("userId is", userId)
		ctx.Send("get user")
	})
	app := kao.New()
	app.Use(router)
	cb := func(port *string) {
		println("Server is listening on port", *port)
	}
	app.Listen(nil, &cb)
}
