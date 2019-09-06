package snowflake

/* snowflake bitmap
 * msb                                                                lsb
 * 0--0000000000-00000000000000000000000000000000-00000000-0000000000000
 * |       |                     |                    |          |
 * unused  round                 time                 machine    serial
*/
const (
	bitAll       = 63
	bitRound     = 10
	bitTime      = 32  // 2^32 / (3600*24*365) = 139.19 year/(1 second/bit)
	bitMachine   = 8
	bitSerial    = bitAll - bitRound - bitTime - bitMachine

	offsetRound  = bitTime + bitMachine + bitSerial
	offsetTime   = bitMachine + bitSerial
	offsetMachine= bitSerial

	maskRound    = uint64((1<<bitRound - 1)   << offsetRound)
	maskTime     = uint64((1<<bitTime - 1)    << offsetTime)
	maskMachine  = uint64((1<<bitMachine - 1) << offsetMachine)
	maskSerial   = uint64(1<<bitSerial - 1)

	RoundMax     = 1<<bitRound - 1
	TimeStart    = 1546300800  // 2019-09-01 00:00 AM
	TimeStop     = TimeStart + 1<<bitTime - 1
	MachineMax   = 1<<bitMachine - 1
	SerialMax    = 1<<bitSerial - 1
)