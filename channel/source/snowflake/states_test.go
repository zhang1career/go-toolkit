package snowflake_test

import (
	"github.com/zhang1career/lib/channel/source/snowflake"
	"testing"
	"time"
)

func TestSnowGroup_Run(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	tic  := time.NewTicker(time.Second)

	req  := make(chan int32)
	defer close(req)

	sg := snowflake.CreateGroup(123)
	got := sg.Run(done, tic, req)

	time.Sleep(time.Second)
	req <- 1
	t.Logf("%0x\r", <-got)

	time.Sleep(time.Second)
	req <- 1
	t.Logf("%0x\r", <-got)

	time.Sleep(time.Second)
	req <- 1
	t.Logf("%0x\r", <-got)
}
