package snowflake_test

import (
	"github.com/zhang1career/lib/channel/source/snowflake"
	"testing"
)

func TestSnowflake_Run(t *testing.T) {
	machine := 0

	var id uint64
	var err error

	snow1 := snowflake.New(machine)

	id, err = snow1.Do(); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	id, err = snow1.Do(); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	id, err = snow1.Do(); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	snow1.Reset()

	id, err = snow1.Do(); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	id, err = snow1.Do(); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	id, err = snow1.Do(); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)


	snow2 := snowflake.New(machine)

	id, err = snow2.Do(); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	id, err = snow2.Do(); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	id, err = snow2.Do(); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	snow2.Reset()

	id, err = snow2.Do(); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	id, err = snow2.Do(); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	id, err = snow2.Do(); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

}
