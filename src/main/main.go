package main

import (
	"log"

	"kao"
	"model"
)

func main() {
	router := kao.NewRouter()
	router.Get("/user/:userId", func(ctx *kao.Ctx) {
		userId, _ := ctx.GetParam("userId")
		ctx.SetCookie("sessionId", "80asd-dsd8-daf988das-88a0")
		ctx.SendMessage(&model.User{
			Name:    userId,
			Age:     999,
			Hobbies: []string{"adsfa", "dddd"},
			Gender:  model.User_FEMALE,
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
