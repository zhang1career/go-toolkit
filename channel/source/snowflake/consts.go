package snowflake

/* snowflake bitmap
 * msb                                                                lsb
 * 0-00000000000000000000000000000000-0000000000-00000000-0000000000000
 * |                  |                   |         |           |
 * reserved           time                round     machine     serial
 * |                  |                   |         |           |
 * +------------------+---------- all ----+---------+-----------+
*/
const (
	bitAll       = 63
	bitTime      = 32  // 2^32 / (3600*24*365) = 139.19 year/(1 second/bit)
	bitRound     = 10
	bitMachine   = 8
	bitSerial    = bitAll - bitTime - bitRound - bitMachine
	
	offsetTime   = bitRound + bitMachine + bitSerial
	offsetRound  = bitMachine + bitSerial
	offsetMachine= bitSerial
	
	maskTime     = uint64((1<<bitTime - 1)    << offsetTime)
	maskRound    = uint64((1<<bitRound - 1)   << offsetRound)
	maskMachine  = uint64((1<<bitMachine - 1) << offsetMachine)
	maskSerial   = uint64(1<<bitSerial - 1)
	
	TimeStart    = 1546300800  // 2019-09-01 00:00 AM
	TimeStop     = TimeStart + 1<<bitTime - 1
	RoundMax     = 1<<bitRound - 1
	MachineMax   = 1<<bitMachine - 1
	SerialMax    = 1<<bitSerial - 1
)