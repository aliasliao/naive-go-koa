package main

import (
	"flag"
	"log"

	"github.com/aliasliao/naive-go-koa/addons/logger"
	"github.com/aliasliao/naive-go-koa/addons/serve"
	"github.com/aliasliao/naive-go-koa/core"
)

var (
	port = flag.Int("port", 3000, "Server port")
	dir  = flag.String("dir", ".", "Public path")
)

func main() {
	flag.Parse()

	app := core.New()
	if s, err := serve.New(*dir); err != nil {
		log.Fatal(err)
	} else {
		app.Use(s)
	}
	app.Use(logger.NewLogger())
	log.Printf("app is listening on %d...", *port)
	if err := app.Listen(*port); err != nil {
		log.Fatal(err)
	}
}
