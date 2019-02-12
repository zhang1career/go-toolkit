package manager

import (
	"github.com/zhang1career/lib/log"
	"sync"
)

type Manager struct {
	Done  chan interface{}
	Wgs   map[string]*sync.WaitGroup
}

func CreateManager() *Manager {
	return &Manager{
		Done: make(chan interface{}),
		Wgs:  make(map[string]*sync.WaitGroup, 0),
	}
}

func (this *Manager) Destroy() {
	close(this.Done)
}


func (this *Manager) CreateWaitGroup(key string) {
	this.Wgs[key] = &sync.WaitGroup{}
}

func (this *Manager) WaitAdd(key string, delta int) {
	
	wg, ok := this.Wgs[key]
	if !ok {
		log.Panic("key %s is not found", key)
		return
	}
	wg.Add(delta)
}

func (this *Manager) WaitDone(key string) {
	wg, ok := this.Wgs[key]
	if !ok {
		log.Panic("key %s is not found", key)
		return
	}
	wg.Done()
}

func (this *Manager) Wait(key string) {
	wg, ok := this.Wgs[key]
	if !ok {
		log.Panic("key %s is not found", key)
		return
	}
	wg.Wait()
}
