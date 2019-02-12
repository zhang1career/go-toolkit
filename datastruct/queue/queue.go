package queue

import "fmt"

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() (value interface{}, err error) {
	if q == nil {
		return 0, fmt.Errorf("Queue nil, Pop failed")
	}
	if len(*q) <= 0 {
		return 0, fmt.Errorf("Queue empty, Pop failed")
	}
	
	head := (*q)[0]
	*q = (*q)[1:]
	return head, nil
}

func (q *Queue) IsEmpty() (isEmpty bool) {
	return len(*q) == 0
}
