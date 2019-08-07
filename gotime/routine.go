package gotime

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

/**
 * Get goroutine id
 */
func Goid() int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}