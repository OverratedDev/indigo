package utils

import (
	"os"
	"strconv"
	"io/ioutil"
	"sync"
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

type FileMutex struct {
	mutex *sync.Mutex
	file  *os.File
}

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

func WriteFile(filepath string, b []byte) error {
	return ioutil.WriteFile(filepath, b, 0660)
}

func GetFile(filepath string) (*os.File, error) {
	return os.OpenFile(filepath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
}

func DeleteFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}
	err = os.Remove(filepath)
	return err
}
