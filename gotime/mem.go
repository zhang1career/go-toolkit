package gotime

import (
	"runtime"
)

type MemInfo struct {
	Alloc    uint64
	Total    uint64
	System   uint64
	NumGc    uint64
}

// return
//   usage
//     Alloc: current memory being used, count in bits
//     Total: total memory being used, count in bits
//     System: OS memory being used, count in bits
//     NumGc: the number of garage collection cycles completed
func MemUsage() (usage MemInfo) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	usage = MemInfo {
		Alloc:  m.Alloc,
		Total:  m.TotalAlloc,
		System: m.Sys,
		NumGc:  uint64(m.NumGC),
	}
	return usage
}
