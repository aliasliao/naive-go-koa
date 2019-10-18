package store

import "time"

type Item struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Done     bool      `json:"done"`
	createAt time.Time `json:"-"`
	updateAt time.Time `json:"-"`
}

type Dao struct {
	Items []Item
}

func New() Dao {
	return Dao{make([]Item, 0)}
}
