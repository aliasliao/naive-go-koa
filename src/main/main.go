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
		ctx.Sendm(&model.User{
			Name:    userId,
			Age:     999,
			Hobbies: []string{"adsfa", "dddd"},
			Gender:  model.User_FEMALE,
		})
	}).Post("/user/:userId", func(ctx *kao.Ctx) {
		userId, _ := ctx.GetParam("userId")
		user := &model.User{
			Name: userId,
		}
		ctx.Parsem(user)
		ctx.Sendm(user)
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

// curl -v -H "Content-Type: application/json;charset=UTF-8" -d '{\"name\":\"123\",\"age\":999,\"hobbies\":[\"adsfa\",\"dddd\"],\"gender\":\"FEMALE\"}' localhost/user/123
