package kv

import (
	"sort"
)

type KV struct {
	Key     interface{}
	Value   interface{}
}

func SortMap(m map[int]interface{}) []KV {
	var keys []int
	
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	
	ret := make([]KV, len(m))
	for i, k := range keys {
		ret[i] = KV{Key: k, Value: m[k]}
	}
	
	return ret
}