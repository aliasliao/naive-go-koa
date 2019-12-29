package main

import (
	"log"

	"github.com/golang/protobuf/proto"
	"kao"
	"model"
)

func main() {
	router := kao.NewRouter()
	router.Get("/user/:userId", func(ctx *kao.Ctx) {
		userId, _ := ctx.GetParam("userId")
		ctx.SetHeader("Content-Type", "application/json;charset=UTF-8")
		ctx.SetCookie("sessionId", "80asd-dsd8-daf988das-88a0")
		user := model.User{
			Name:    userId,
			Age:     999,
			Hobbies: []string{"adsfa", "dddd"},
		}
		ctx.Send(proto.Marshal(&user))
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
