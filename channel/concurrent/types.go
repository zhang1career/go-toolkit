package concurrent

import "github.com/zhang1career/lib/channel/ctrlbus"

type Work interface {
	Do(interface{}) (interface{}, error)
}

type parser struct {
	managers    []*ctrlbus.Ctrlbus
	wholesaler  *wholesaler
	tasks       chan *task
	products    chan *product
}

type wholesaler struct {
	teams       chan *team
}

type team struct {
	id      int
	broker  chan *task
	worker  *worker
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
	err     error
}