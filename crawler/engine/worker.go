package engine

import (
	"github.com/zhang1career/golab/io/net"
	"github.com/zhang1career/golab/log"
)

type worker struct {
	dumps map[string]bool
}

func (w *worker) explore(node interface{}) (interface{}, error) {
	r := index(node)
	body, err := net.Curl(r.Url)
	if err != nil {
		log.Error("fail to fetch url[%s]: %v", r.Url, err.Error())
		return nil, err
	}
	log.Info("fetch url[%s]...", r.Url)
	
	return ExploreResult{Parse:r.Parse, Data:body}, nil
}

func (w *worker) parse(r interface{}) ([]interface{}, interface{}) {
	children := make([]interface{}, 0)
	parseResult := r.(ExploreResult).Parse(r.(ExploreResult).Data)
	for _, rr := range parseResult.Requests {
		children = append(children, rr)
	}
	
	return children, parseResult.Items
}

func (w *worker) root(node interface{}) bool {
	r := index(node)
	if w.dumps[r.Url] {
		return true
	}
	w.dumps[r.Url] = true
	return false
}

func (w *worker) done(node interface{}) (done bool) {
	return false
}
