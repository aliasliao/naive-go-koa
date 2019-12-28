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
		ctx.SetHeader("Content-Type", "application/json;charset=UTF-8")
		ctx.SetCookie("sessionId", "80asd-dsd8-daf988das-88a0")
		ctx.Send(User{
			id:   userId,
			name: "jack",
			age:  99,
		})
	})
	app := kao.New()
	app.Use(router)
	err := app.Listen(80, func(port *string) {
		log.Println("Server is listening on port", *port)
	})
	if err != nil {
		log.Fatal(err)
	}
}
