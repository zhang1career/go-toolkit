package cache

import (
	"sync"
)

var once sync.Once
var instance *Cache

func NewOnce(url string, port int, ext map[string]interface{}) *Cache {
	once.Do(func() {
		instance = &Cache{
			conn: New(url, port, ext),
		}
	})
	return instance
}