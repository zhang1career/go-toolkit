package snowflake_test

import (
	"github.com/zhang1career/lib/channel/source/snowflake"
	"math/rand"
	"testing"
)

func TestCompose_Decompose(t *testing.T) {
	for r := 0; r < 1023; r++ {
		for m := 0; m <= snowflake.MachineMax; m++ {
			g := snowflake.GetGroupBitcode(r, snowflake.GetTime(), m)
			s := rand.Intn(snowflake.SerialMax + 1)
			z := snowflake.SnowMake(g, s)
			round, _, machine, serial := snowflake.SnowMelt(z)
			if r != int(round) {
				t.Logf("error of r, expect %d, got %d", r, round)
				continue
			}
			if m != int(machine) {
				t.Logf("error of m, expect %d, got %d", m, machine)
				continue
			}
			if s != int(serial) {
				t.Logf("error of s, expect %d, got %d", s, serial)
				continue
			}
		}
	}
}