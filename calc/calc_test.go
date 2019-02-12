package calc

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func swap(a, b *int){
	a, b = b, a
}

func source(count int) <-chan int {
	out := make(chan int)
	go func() {
		rand.Seed(time.Now().Unix())
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

func TestApply(t *testing.T) {
	a, b := 3,4
	swap(&a, &b)
	fmt.Println(a, b)
	
	result, err := Apply(Mul, 13, 4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}