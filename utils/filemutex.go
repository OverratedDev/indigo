package utils

import (
	"os"
	"sync"
)

type FileMutex struct {
	mutex *sync.Mutex
	file  *os.File
}