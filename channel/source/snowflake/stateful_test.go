package snowflake_test

import (
	"github.com/zhang1career/lib/channel/source/snowflake"
	"testing"
)

func TestSnowflake_Do_Undo(t *testing.T) {
	var salt interface{}
	salt = uint64(0)
	for i := 0; i < 5000; i++ {
		worker := snowflake.CreateSnowFlake(i)
		salt = worker.Do(salt)
		timestamp, machine, serial := worker.Undo(salt)
		if i != machine {
			t.Logf("[OVERFLOW] i=%d, t=%d, m=%d, s=%d", i, timestamp, machine, serial)
			continue
		}
		t.Logf("i=%d, t=%d, m=%d, s=%d", i, timestamp, machine, serial)
	}
}
