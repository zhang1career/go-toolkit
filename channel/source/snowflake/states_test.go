package snowflake_test

import (
	"github.com/zhang1career/lib/channel/ctrlbus"
	"github.com/zhang1career/lib/channel/source/snowflake"
	"testing"
	"time"
)

func TestSnowGroup_Run(t *testing.T) {
	config := map[string]interface{}{
		"ticker": time.Second,
	}
	manager := ctrlbus.CreateCtrlbus(config)

	req  := make(chan int32)
	defer close(req)

	sg := snowflake.CreateGroup(123)
	got := sg.Run(manager, req)

	time.Sleep(time.Second)
	req <- 1
	t.Logf("%0x\r", <-got)

	time.Sleep(time.Second)
	req <- 1
	t.Logf("%0x\r", <-got)
	req <- 1
	t.Logf("%0x\r", <-got)

	time.Sleep(time.Second)
	req <- 1
	t.Logf("%0x\r", <-got)
	req <- 1
	t.Logf("%0x\r", <-got)
	req <- 1
	t.Logf("%0x\r", <-got)
}
