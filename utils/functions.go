package utils

import (
	"strconv"
)

const (
	MINUTE = 60
	HOUR   = MINUTE * 60
	DAY    = HOUR * 24
	MONTH  = DAY * 30
	YEAR   = MONTH * 12
	BYTE   = 1
	KB     = 1024 * BYTE
	MB     = 1024 * KB
	GB     = 1024 * MB
)

func FormatUptime(startTime, currentTime int64) string {
	val := currentTime - startTime
	if val < MINUTE {
	 	return strconv.Itoa(int(val)) + "s"
	} else if val >= MINUTE && val < HOUR {
		return strconv.Itoa(int(val/MINUTE)) + "m"
	} else if val >= HOUR && val < DAY {
		return strconv.Itoa(int(val/HOUR)) + "h"
	} else if val >= DAY && val < MONTH {
		return strconv.Itoa(int(val/DAY)) + "d"
	} else if val >= MONTH && val < YEAR {
		return strconv.Itoa(int(val/MONTH)) + "M"
	}
	return strconv.Itoa(int(val/YEAR)) + "y"
}