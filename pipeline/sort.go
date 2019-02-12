package pipeline

import (
	"sort"
)


func InMemSort(in <-chan int) <-chan int {
	out := make(chan int, ChanBuffSize)
	go func() {
		// read into memory
		memory := []int{}
		for v := range in {
			memory = append(memory, v)
		}
		// sort
		sort.Ints(memory)
		// output
		for _, v := range memory {
			out <- v
		}
		close(out)
	}()
	return out
}
