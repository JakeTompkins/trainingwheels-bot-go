package database

import (
	"sync"

	c "github.com/ostafen/clover"
)

var client *c.DB

func GetDb() *c.DB {
	var once sync.Once
	onceBody := func() {
		client, _ = c.Open("clover-db")
	}

	once.Do(onceBody)

	return client
}
