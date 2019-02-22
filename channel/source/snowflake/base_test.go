package snowflake_test

import (
	"github.com/zhang1career/lib/channel/source/snowflake"
	"testing"
)

func TestCompose_Decompose(t *testing.T) {
	for i := 0; i < 5000; i++ {
		timestamp, machine, serial := snowflake.SnowMelt(snowflake.SnowMake(uint64(i)))
		if i != machine {
			t.Logf("[OVERFLOW] i=%d, t=%d, m=%d, s=%d", i, timestamp, machine, serial)
			continue
		}
	}
}