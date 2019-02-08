package watcher

import (
	"os"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/vereas/indigo/process"
)

type PStatus struct {
	state *os.ProcessState
	err   error
}

type ProcWatcher struct {
	procStatus chan *PStatus
	proc 	   process.ProcContainer
	stopWatcher chan bool
}

type Watcher struct {
	sync.Mutex
	restartProc chan ProcContainer
	watchProcs  map[string]*ProcWatcher
}

func InitWatcher() *Watcher {
	watcher := &Watcher{
		restartProc: make(chan ProcContainer),
		watchProcs:  make(map[string]*ProcWatcher),
	}
	return watcher
}

func (watcher *Watcher) RestartProc() chan ProcContainer {
	return watcher.restartProc
}

func (watcher *Watcher) AddProcWatcher(proc ProcContainer) {
	watcher.Lock()
	defer watcher.Unlock()
	if _, ok := watcher.watchProcs[proc.Identifier()]; ok {
		log.Warnf("A watcher for this process already exists.")
		return
	}
	procWatcher := &ProcWatcher {
		procStatus:  make(chan *PStatus, 1),
		proc:        proc,
		stopWatcher: make(chan bool, 1),
	}
	watcher.watchProcs[proc.Identifier()] = procWatcher
	go func() {
		log.Infof("Starting watcher on proc %s", proc.Identifier)
		state, err := proc.Watch()
		procWatcher.procStatus <- &PStatus{
			state: state,
			err:   err,
		}
	}()
	go func() {
		defer delete(watcher.watchProcs, procWatcher.proc.Identifier())
		select {
		case procStatus := <-procWatcher.procStatus:
			log.Infof("Proc %s is dead, advising master...", procWatcher.proc.Identifier())
			log.Infof("State is %s", procStatus.state.String())
			watcher.restartProc <- procWatcher.proc
			break
		case <-procWatcher.stopWatcher:
			break
		}
	}()
}