package main

import (
	"sync"

	"github.com/vereas/indigo/process"
)

type Master struct {
	sync.Mutex

	SysFolder string
	PidFile   string
	OutFile   string
	ErrFile   string
	Watcher   *Watcher

	Procs map[string]process.ProcContainer
}