package source

import (
	"math/rand"
	"time"
)

func RandSource(count int) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		rand.Seed(time.Now().Unix())
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
	}()
	return out
}
