package main

import (
	"log"

	"kao"
)

type User struct {
	id   string
	name string
	age  int
}

func main() {
	router := kao.NewRouter()
	router.Get("/user/:userId", func(ctx *kao.Ctx) {
		userId, _ := ctx.GetParam("userId")
		ctx.Send(User{
			id:   userId,
			name: "jack",
			age:  99,
		})
	})
	app := kao.New()
	app.Use(router)
	cb := func(port *string) {
		log.Println("Server is listening on port", *port)
	}
	err := app.Listen(nil, &cb)
	if err != nil {
		log.Fatal(err)
	}
}
