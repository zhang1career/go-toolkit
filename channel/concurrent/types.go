package concurrent

import "github.com/zhang1career/lib/channel/ctrlbus"

type Work interface {
	Do(interface{}) interface{}
}

type parser struct {
	managers    []*ctrlbus.Ctrlbus
	wholesaler  chan *team
	products    chan *product
}

type team struct {
	id      int
	broker  chan *task
	worker  worker
}

type worker struct {
	Work
}

type task struct {
	id      int
	job     interface{}
}

type product struct {
	taskId  int
	output  interface{}
}