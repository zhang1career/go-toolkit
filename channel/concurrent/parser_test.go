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

func (t *testWorker) Do(in interface{}) (interface{}, error) {
	return t.id + in.(int), nil
}


func TestParser_Parse(t *testing.T) {
	parser := concurrent.CreateParser()

	config := map[string]interface{}{
		"ticker": time.Second,
	}
	parser.AddTeam(config, 3, createTester)

	parser.Run()

	ret, err := parser.Parse(1); if err != nil {
		t.Error(err.Error())
	}
	t.Log(ret)
}

