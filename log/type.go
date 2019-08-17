package log

const (
	lvDebug int = iota + 0
	lvTrace
	lvInfo
	lvWarn
	lvError
	lvFatal
)

var LvMsg = map[int]string {
	lvDebug : "DEBUG",
	lvTrace : "TRACE",
	lvInfo  : "INFO",
	lvWarn  : "WARN",
	lvError : "ERROR",
	lvFatal : "FATAL",
}