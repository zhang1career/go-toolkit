package bit

func Byte2KB(b uint64) uint64 {
	return b >> 10
}

func Byte2MB(b uint64) uint64 {
	return b >> 20
}
