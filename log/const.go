package log

const (
	lvDebug int = iota + 0
	lvInfo
	lvWarn
	lvError
)

var LvMsg = map[int]string {
	lvDebug : "[DEBUG]",
	lvInfo  : "[INFO]",
	lvWarn  : "[WARN]",
	lvError : "[ERROR]",
}