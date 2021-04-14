package concurrent_test

import (
	"github.com/zhang1career/golab/channel/concurrent"
	"github.com/zhang1career/golab/channel/ctrlbus"
	"testing"
	"time"
)

type testWorker struct {
	id  int
}

func createTester(ctrlbus *ctrlbus.Ctrlbus, id int) concurrent.Work {
	return &testWorker{id: id}
}

func (t *testWorker) Do(in interface{}) concurrent.Output {
	return concurrent.CreateOutput(t.id + in.(int), nil)
}


func TestParser_Parse(t *testing.T) {
	parser := concurrent.CreateParser()

	config := map[string]interface{}{
		"ticker": time.Second,
	}
	parser.AddWorker(config, 32, createTester)

	parser.Run()

	t.Log(parser.Parse(0))

	sum := make(map[int]int, 32)

	for i := 0; i < 100000; i++ {
		output := parser.Parse(0)
		sum[output.GetValue().(int)]++
	}

	t.Log(sum)
}

