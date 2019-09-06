package snowflake_test

import (
	"github.com/zhang1career/lib/channel/source/snowflake"
	"testing"
)

func TestSnowGroup_Reset_And_Do(t *testing.T) {
	machine := 0

	var id []uint64
	var err error

	sg1 := snowflake.CreateGroup(machine)

	id, err = sg1.Do(1); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	id, err = sg1.Do(1); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	id, err = sg1.Do(1); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	sg1.Reset()

	id, err = sg1.Do(1); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	id, err = sg1.Do(1); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	id, err = sg1.Do(1); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)


	sg2 := snowflake.CreateGroup(machine)

	id, err = sg2.Do(1); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	id, err = sg2.Do(1); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	id, err = sg2.Do(1); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	sg2.Reset()

	id, err = sg2.Do(1); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	id, err = sg2.Do(1); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

	id, err = sg2.Do(1); if err != nil {
		t.Error(err.Error())
	}
	t.Logf("%0x\r", id)

}
