package log

import (
	"fmt"
	"github.com/zhang1career/lib/datastruct/bit"
	"github.com/zhang1career/lib/gotime"
	"time"
)


func Debug(message string, args ...interface{}) {
	msg := check(message, args...)
	before(lvDebug)
	fmt.Printf("%s\n", msg)
	after(lvDebug)
}

func Info(message string, args ...interface{}) {
	msg := check(message, args...)
	before(lvInfo)
	fmt.Printf("%s\n", msg)
	after(lvInfo)
}

func Warn(message string, args ...interface{}) {
	msg := check(message, args...)
	before(lvWarn)
	fmt.Printf("%s\n", msg)
	after(lvWarn)
}

func Error(message string, args ...interface{}) {
	msg := check(message, args...)
	before(lvError)
	fmt.Printf("%s\n", msg)
	after(lvError)
}

func Fatal(message string, args ...interface{}) {
	msg := check(message, args...)
	before(lvFatal)
	fmt.Printf("%s\n", msg)
	after(lvFatal)
}

func Panic(message string, args ...interface{}) {
	where := here(4)
	msg := check(message, args...)
	panic(where + msg)
}

func check(message string, args ...interface{}) (msg string) {
	msg = message
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	return msg
}

func before(level int) {
	t := time.Now().Format("2006-01-02 15:04:05")
	lv := LvMsg[level]
	fmt.Printf("%s %s ", t, lv)
	
	if level > lvWarn {
		fmt.Printf("%s\n", here(4))
	}
}

func after(level int) {
	if level < lvInfo {
		mem()
	}
}

func here(skip int) string {
	fun, file, line, err := gotime.Locate(skip)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s@%s:%d ", fun, file, line)
}

func mem() {
	m := gotime.MemUsage()
	fmt.Printf("MemoryInfo\t\theap:%dMB\ttotal_heap:%dMB\tos:%dMB\tgc:%d\n\n",
		bit.Byte2MB(m.Alloc),
		bit.Byte2MB(m.Total),
		bit.Byte2MB(m.System),
		bit.Byte2MB(m.NumGc))
}
