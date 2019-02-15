package ctrlbus

type Ctrlbus struct {
	Done  chan interface{}
}

func CreateCtrlbus() *Ctrlbus {
	return &Ctrlbus{
		Done: make(chan interface{}),
	}
}

func (this *Ctrlbus) Destroy() {
	close(this.Done)
}
