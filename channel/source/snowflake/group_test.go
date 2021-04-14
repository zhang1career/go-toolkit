package snowflake_test

import (
	"github.com/zhang1career/golab/channel/concurrent"
	"github.com/zhang1career/golab/channel/ctrlbus"
	"github.com/zhang1career/golab/channel/source/snowflake"
	"testing"
	"time"
)

func TestSnowGroup_Reset_And_Do(t *testing.T) {
	machine := 0

	var output concurrent.Output

	config := map[string]interface{}{
		"ticker": time.Second,
	}
	ctrlbus := ctrlbus.CreateCtrlbus(config)

	sg1 := snowflake.CreateGroup(ctrlbus, machine)

	output = sg1.Do(1); if output.GetErr() != nil {
		t.Error(output.GetErr().Error())
	}
	t.Logf("%0x\r", output.GetValue())

	output = sg1.Do(1); if output.GetErr() != nil {
		t.Error(output.GetErr().Error())
	}
	t.Logf("%0x\r", output.GetValue())

	output = sg1.Do(1); if output.GetErr() != nil {
		t.Error(output.GetErr().Error())
	}
	t.Logf("%0x\r", output.GetValue())

	sg1.Reset()

	output = sg1.Do(1); if output.GetErr() != nil {
		t.Error(output.GetErr().Error())
	}
	t.Logf("%0x\r", output.GetValue())

	output = sg1.Do(1); if output.GetErr() != nil {
		t.Error(output.GetErr().Error())
	}
	t.Logf("%0x\r", output.GetValue())

	output = sg1.Do(1); if output.GetErr() != nil {
		t.Error(output.GetErr().Error())
	}
	t.Logf("%0x\r", output.GetValue())


	sg2 := snowflake.CreateGroup(ctrlbus, machine)

	output = sg2.Do(1); if output.GetErr() != nil {
		t.Error(output.GetErr().Error())
	}
	t.Logf("%0x\r", output.GetValue())

	output = sg2.Do(1); if output.GetErr() != nil {
		t.Error(output.GetErr().Error())
	}
	t.Logf("%0x\r", output.GetValue())

	output = sg2.Do(1); if output.GetErr() != nil {
		t.Error(output.GetErr().Error())
	}
	t.Logf("%0x\r", output.GetValue())

	sg2.Reset()

	output = sg2.Do(1); if output.GetErr() != nil {
		t.Error(output.GetErr().Error())
	}
	t.Logf("%0x\r", output.GetValue())

	output = sg2.Do(1); if output.GetErr() != nil {
		t.Error(output.GetErr().Error())
	}
	t.Logf("%0x\r", output.GetValue())

	output = sg2.Do(1); if output.GetErr() != nil {
		t.Error(output.GetErr().Error())
	}
	t.Logf("%0x\r", output.GetValue())
}
