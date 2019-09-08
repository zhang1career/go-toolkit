package snowflake_test

import (
	"github.com/zhang1career/lib/channel/source/snowflake"
	"testing"
	"time"
)


func TestSnowQueue_GetId(t *testing.T) {
	config := map[string]interface{}{
		"ticker": time.Second,
	}
	q := snowflake.CreateQueue(config, 256)

	for i := 0; i < 8193; i++ {
		id, err := q.GetId()
		if err != nil {
			t.Log(err.Error())
			continue
		}
		t.Log(id)
	}
}

func BenchmarkSnowQueue_GetId(b *testing.B) {
	b.StopTimer()

	config := map[string]interface{}{
		"ticker": time.Second,
	}
	q := snowflake.CreateQueue(config, 256)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, err := q.GetId()
		if err != nil {
			b.Log(err.Error())
			continue
		}
		//b.Log(id)
	}
}
