package headers

//Location: include/uapi/linux/time.h
var ClockLookup = map[int]string{
	0:  "CLOCK_REALTIME",
	1:  "CLOCK_MONOTONIC",
	2:  "CLOCK_PROCESS_CPUTIME_ID",
	3:  "CLOCK_THREAD_CPUTIME_ID",
	4:  "CLOCK_MONOTONIC_RAW",
	5:  "CLOCK_REALTIME_COARSE",
	6:  "CLOCK_MONOTONIC_COARSE",
	7:  "CLOCK_BOOTTIME",
	8:  "CLOCK_REALTIME_ALARM",
	9:  "CLOCK_BOOTTIME_ALARM",
	10: "CLOCK_SGI_CYCLE",
	11: "CLOCK_TAI",
}