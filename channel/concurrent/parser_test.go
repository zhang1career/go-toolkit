package concurrent_test

import (
	"github.com/zhang1career/lib/channel/concurrent"
	"testing"
	"time"
)

type testWorker struct {
	id  int
}

func createTester(id int) concurrent.Work {
	return &testWorker{id: id}
}

func (t *testWorker) Do(in interface{}) interface{} {
	return t.id + in.(int)
}


func TestParser_Parse(t *testing.T) {
	parser := concurrent.CreateParser()

	config := map[string]interface{}{
		"ticker": time.Second,
	}
	parser.AddTeam(config, 32, createTester)

	parser.Run()

	t.Log(parser.Parse(0))

	sum := make(map[int]int, 32)

	for i := 0; i < 100000; i++ {
		index := parser.Parse(0).(int)
		sum[index]++
	}

	t.Log(sum)
}

