package stack_test

import (
	"github.com/zhang1career/golab/datastruct/stack"
	"testing"
)

func TestStack(t *testing.T) {
	datas := []struct{in, out int} {
		{0, 0},
		{1, 1},
	}
	
	for _, data := range datas {
		s := stack.New()
		s.Push(data.in)
		actual := s.Pop()
		if actual != data.out {
			t.Errorf("stack.Push(%d), stack.Pop() %d, expected %d", data.in, actual, data.out)
		}
	}
}

func BenchmarkStack(b *testing.B) {
	data := struct{in, out int} {0, 0}
	
	for i := 0; i < b.N; i++ {
		s := stack.New()
		s.Push(data.in)
		actual := s.Pop()
		if actual != data.out {
			b.Errorf("queue.Push(%d), queue.Pop() %d, expected %d", data.in, actual, data.out)
		}
	}
}