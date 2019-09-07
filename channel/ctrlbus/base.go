package ctrlbus

import "time"

type Ctrlbus struct {
	done    chan interface{}
	ticker  *time.Ticker
}

func CreateCtrlbus(ext map[string]interface{}) *Ctrlbus {
	ret := &Ctrlbus{
		done:   make(chan interface{}),
	}
	ticker, ok := ext["ticker"]; if ok {
		ret.ticker = time.NewTicker(ticker.(time.Duration))
	}
	return ret
}

// done
func (c *Ctrlbus) Destroy() {
	close(c.done)
}

func (c *Ctrlbus) GetDone() chan interface{} {
	return c.done
}

// ticker
func (c *Ctrlbus) SetTicker(t time.Duration) {
	c.ticker = time.NewTicker(t)
}

func (c *Ctrlbus) GetTicker() *time.Ticker {
	return c.ticker
}