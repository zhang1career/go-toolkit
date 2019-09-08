package snowflake_test

import (
	"github.com/zhang1career/lib/channel/source/snowflake"
	"testing"
	"time"
)


func TestGetId(t *testing.T) {
	config := map[string]interface{}{
		"ticker": time.Second,
	}
	q := snowflake.CreateQueue(config)

	for i := 0; i < 8193; i++ {
		id, err := q.GetId()
		if err != nil {
			t.Log(err.Error())
			continue
		}
		t.Log(id)
	}
}
