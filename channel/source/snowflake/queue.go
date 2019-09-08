package snowflake

import (
	"github.com/zhang1career/lib/channel/concurrent"
)

func CreateQueue(config map[string]interface{}, count int) *SnowQueue {
	//
	p := concurrent.CreateParser()
	//
	p.AddTeam(config, count, CreateGroupAsWorker)
	//
	p.Run()

	return &SnowQueue{parser: p}
}

func (q *SnowQueue) GetId() ([]uint64, error) {
	out := q.parser.Parse(1)
	if out.GetErr() != nil {
		return nil, out.GetErr()
	}
	return out.GetValue().([]uint64), nil
}