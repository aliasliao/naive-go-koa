package main

import (
	"log"
	"net/http"
	"time"
	"utils"
)

type Item struct {
	id       string
	title    string
	done     bool
	createAt time.Time
	updateAt time.Time
}

type Dao struct {
	items  []Item
	routes *router.Mux
}

func (dao *Dao) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method, url := r.Method, r.URL.String()
	fallbackAction := (*dao.routes)["*"]["*"]
	if dict, exist1 := (*dao.routes)[method]; exist1 {
		if action, exist2 := dict[url]; exist2 {
			action(w, r)
		} else {
			fallbackAction(w, r)
		}
	} else {
		fallbackAction(w, r)
	}
}

func newDao(route *router.Route) *Dao {
	return &Dao{
		items:  make([]Item, 0),
		routes: &((*route).Routes),
	}
}

func main() {
	route := router.New()
	route.Get("/api/items", func(w http.ResponseWriter, r *http.Request) {

	}).Post("/api/item", func(w http.ResponseWriter, r *http.Request) {

	}).Patch("/api/item/:id", func(w http.ResponseWriter, r *http.Request) {

	}).Delete("/api/item/:id", func(w http.ResponseWriter, r *http.Request) {

	})

	http.Handle("/", newDao(route))
	if e := http.ListenAndServe(":8080", nil); e != nil {
		log.Fatal(e)
	}
}
