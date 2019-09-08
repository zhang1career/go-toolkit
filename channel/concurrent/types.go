package concurrent

import "github.com/zhang1career/lib/channel/ctrlbus"

type Work interface {
	Do(interface{}) interface{}
}

type parser struct {
	managers    []*ctrlbus.Ctrlbus
	seekers     chan chan interface{}
	input       chan interface{}
	output      chan interface{}
}

type worker struct {
	id      int
	Work
}