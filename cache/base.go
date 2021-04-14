package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/zhang1career/golab/log"
	"time"
)

func New(url string, port int, ext map[string]interface{}) *redis.Client {
	pass, ok := ext["pass"].(string); if !ok {
		pass = ""
	}
	db, ok := ext["db"].(int); if !ok {
		db = 0
	}

	ret := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", url, port),
		Password: pass,
		DB:       db,
	})
	if ret == nil {
		log.Fatal("fail to create redis client")
	}
	return ret
}

func (this *Cache) Get(key string) (string, bool) {
	val, err := this.conn.Get(key).Result()
	if err != nil {
		log.Error("fail to get key[%s]", key)
	}
	return val, err == nil
}

func (this *Cache) Set(key string, value string, ext map[string]interface{}) bool {
	expire, ok := ext["expire"].(time.Duration); if !ok {
		expire = 0
	}

	err := this.conn.Set(key, value, expire).Err()
	if err != nil {
		log.Error("fail to set key[%s], value[%s]", key, value)
	}
	return err == nil
}

func (this *Cache) Incr(key string) (int64, bool) {
	val, err := this.conn.Incr(key).Result()
	if err != nil {
		log.Error("fail to incr key[%s]", key)
	}
	return val, err == nil
}
