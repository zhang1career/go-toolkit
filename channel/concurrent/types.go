package concurrent

import "github.com/zhang1career/golab/channel/ctrlbus"

type Work interface {
	Do(interface{}) Output
}

type Parser struct {
	managers    []*ctrlbus.Ctrlbus
	seekers     chan chan interface{}
	input       chan interface{}
	output      chan Output
}

type worker struct {
	id      int
	Work
}

type Output struct {
	value       interface{}
	err         error
}