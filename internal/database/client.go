package database

import (
	"sync"

	c "github.com/ostafen/clover"
)

var client *c.DB

var (
	once     sync.Once
	onceBody = func() {
		client, _ = c.Open("clover-db")
	}
)

func GetDb() *c.DB {
	once.Do(onceBody)

	return client
}
