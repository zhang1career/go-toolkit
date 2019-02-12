package queue

import "testing"

func TestQueue(t *testing.T) {
	tests := []struct{in, out int} {
		{0, 0},
		{1, 1},
	}
	
	for _, tt := range tests {
		q := Queue{}
		q.Push(tt.in)
		actual, err := q.Pop()
		if err != nil {
			t.Errorf("queue.Push(%d), queue.Pop() error: %s", tt.in, err.Error())
		} else if actual != tt.out {
			t.Errorf("queue.Push(%d), queue.Pop() %d, expected %d", tt.in, actual, tt.out)
		}
	}
}

func BenchmarkQueue(b *testing.B) {
	tt := struct{in, out int} {0, 0}
	
	for i := 0; i < b.N; i++ {
		q := Queue{}
		q.Push(tt.in)
		actual, err := q.Pop()
		if err != nil {
			b.Errorf("queue.Push(%d), queue.Pop() error: %s", tt.in, err.Error())
		} else if actual != tt.out {
			b.Errorf("queue.Push(%d), queue.Pop() %d, expected %d", tt.in, actual, tt.out)
		}
	}
}