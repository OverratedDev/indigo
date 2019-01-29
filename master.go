package main

import (
	"sync"
)

type Master struct {
	sync.Mutex

	SysFolder string
	PidFile   string
	OutFile   string
	ErrFile   string
	Watcher   *Watcher

	Procs map[string]ProcContainer
}