package gotime

import (
	"github.com/zhang1career/golab/datastruct/bit"
	"runtime"
	"testing"
	"time"
)

func TestMemUsage(t *testing.T) {
	// Print our starting memory usage (should be around 0MB)
	output(t, MemUsage())
	
	var overall [][]int
	for i := 0; i<4; i++ {
		// Allocate memory using make() and append to overall (so it doesn't get
		// garbage collected). This is to create an ever increasing memory usage
		// which we can track. We're just using []int as an example.
		a := make([]int, 0, 999999)
		overall = append(overall, a)
		
		// Print our memory usage at each interval
		output(t, MemUsage())
		time.Sleep(time.Second)
	}
	
	// Clear our memory and print usage, unless the GC has run 'Alloc' will remain the same
	overall = nil
	output(t, MemUsage())
	
	// Force GC to clear up, should see a memory drop
	runtime.GC()
	output(t, MemUsage())
}

func output(t *testing.T, m MemInfo) {
	t.Logf("MemoryInfo\t\theap:%dMB\ttotal_heap:%dMB\tos:%dMB\tgc:%d\n",
		bit.Byte2MB(m.Alloc),
		bit.Byte2MB(m.Total),
		bit.Byte2MB(m.System),
		bit.Byte2MB(m.NumGc))
}
