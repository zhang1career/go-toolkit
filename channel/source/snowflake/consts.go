package snowflake

/* snowflake bitmap
 * msb                                                               lsb
 * 0-000000000000000000000000000000000000000-0000000000000000-00000000
 * |                    |                            |            |
 * reserved             createTime(t)                machineId(m) serialNo(s)
 *                      |                            |            |
 *                      +------------------ all -----+------------+
*/
const (
	bitAll       = 63
	bitT         = 39  // 2^39 / (1000*3600*24*365) = 17.4 year/(1 millisecond/bit)
	bitM         = bitAll - bitT - bitS
	bitS         = 12
	
	maskT        = uint64((1<<bitT - 1) << (bitM + bitS))
	maskM        = uint64((1<<bitM - 1) << bitS)
	maskS        = uint64(1<<bitS - 1)
	
	bitNano2Mil  = 23  // 10'000'000'000 nanosecond >> 23 = 1 ms
)