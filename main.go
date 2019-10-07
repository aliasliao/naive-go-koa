package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)
import (
	"dao"
	"router"
)

func main() {
	route := router.New()
	store := dao.New()

	route.Get("/api/items", func(w http.ResponseWriter, r *http.Request) {
		bytes, e := json.Marshal(store.Items)
		if e != nil {
			log.Printf("Error stringify items: %v", e)
			http.Error(w, "can't stringify items", http.StatusInternalServerError)
			return
		}
		if _, e := io.WriteString(w, string(bytes)); e != nil {
			log.Printf("Error sending items: %v", e)
			http.Error(w, "can't sending items", http.StatusInternalServerError)
			return
		}
	}).Post("/api/item", func(w http.ResponseWriter, r *http.Request) {
		body, e := ioutil.ReadAll(r.Body)
		if e != nil {
			log.Printf("Error reading body: %v", e)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}
		var newItem dao.Item
		if e := json.Unmarshal(body, &newItem); e != nil {
			log.Printf("Error parsing body to json: %v", e)
			http.Error(w, "body is invalid item json", http.StatusBadRequest)
			return
		}
		store.Items = append(store.Items, newItem)
	}).Patch("/api/item/:id", func(w http.ResponseWriter, r *http.Request) {

	}).Delete("/api/item/:id", func(w http.ResponseWriter, r *http.Request) {

	})

	http.Handle("/", route)
	if e := http.ListenAndServe(":8080", nil); e != nil {
		log.Fatal(e)
	}
}
