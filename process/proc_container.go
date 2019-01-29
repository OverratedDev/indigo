package main

import (
	"errors"
	"os"
	"strconv"
	"syscall"

	"github.com/vereas/indigo/utils"
)

type ProcContainer interface {
	Start() error
	ForceStop() error
	GracefullyStop() error
	Restart() error
	Delete() error
	IsAlive() bool
	Identifier() string
	ShouldKeepAlive() bool
	AddRestart()
	NotifyStopped()
	SetStatus(status string)
	SetUptime()
	SetSysInfo()
	GetPid() int
	GetStatus() *ProcStatus
	Watch() (*os.ProcessState, error)
	release()
	GetOutFile() string
	GetPidFile() string
	GetPath() string
	GetErrFile() string
	GetName() string
} 

type Proc struct {
	Name	  string
	Cmd       string
	Args      []string
	Path      string
	Pidfile   string
	Outfile   string
	Errfile   string
	KeepAlive bool
	Pid       int
	Status    *ProcStatus
	process   *os.Process
}

func (proc *Proc) Start() error {
	outFile, err := utils.GetFile(proc.Outfile)
	if err != nil {
		return err
	}
	errFile, err := utils.GetFile(proc.Errfile)
	if err != nil {
		return err
	}
	wd, _ := os.Getwd()
	procAtr := &os.ProcAttr{
		Dir: wd,
		Env: os.Environ(),
		Files: []*os.File{
			os.Stdin,
			outFile,
			errFile,
		},
	}
	args := append([]string{proc.Name}, proc.Args...)
	process, err := os.StartProcess(proc.Cmd, args, procAtr)
	if err != nil {
		return err
	}
	proc.process = process
	proc.Pid = proc.process.Pid
	err = utils.WriteFile(proc.Pidfile, []byte(strconv.Itoa(proc.process.Pid)))
	if err != nil {
		return err
	}
	proc.Status.InitUptime()
	proc.Status.SetStatus("started")
	return nil
}

func (proc *Proc) ForceStop() error {
	if proc.process != nil {
		err := proc.process.Signal(syscall.SIGKILL)
		proc.Status.SetStatus("stopped")
		proc.release()
		return err
	}
	return errors.New("Process does not exist.")
}

func (proc *Proc) GracefullyStop() error {
	if proc.process != nil {
		err := proc.process.Signal(syscall.SIGTERM)
		proc.Status.SetStatus("asked to stop")
		return err
	}
	return errors.New("Process does not exist.")
}

func (proc *Proc) Restart() error {
	if proc.IsAlive() {
		err := proc.GracefullyStop()
		if err != nil {
			return err
		}
	}
	return proc.Start()
}

func (proc *Proc) Delete() error {
	proc.release()
	err := utils.DeleteFile(proc.Outfile)
	if err != nil {
		return err
	}
	err = utils.DeleteFile(proc.Errfile)
	if err != nil {
		return err
	}
	return os.RemoveAll(proc.Path)
}

func (proc *Proc) IsAlive() bool {
	p, err := os.FindProcess(proc.Pid)
	if err != nil {
		return false
	}
	return p.Signal(syscall.Signal(0)) == nil
}

func (proc *Proc) Watch() (*os.ProcessState, error) {
	return proc.process.Wait()
}

func (proc *Proc) release() {
	if proc.process != nil {
		proc.process.Release()
	}
	utils.DeleteFile(proc.Pidfile)
}

func (proc *Proc) NotifyStopped() {
	proc.Pid = -1
}

func (proc *Proc) AddRestart() {
	proc.Status.AddRestart()
}

func (proc *Proc) GetPid() int {
	return proc.Pid
}

func (proc *Proc) GetOutFile() string {
	return proc.Outfile
}

func (proc *Proc) GetErrFile() string {
	return proc.Errfile
}

func (proc *Proc) GetPidFile() string {
	return proc.Pidfile
}

func (proc *Proc) GetPath() string {
	return proc.Path
}

func (proc *Proc) GetStatus() *ProcStatus {
	if !proc.IsAlive() {
		proc.ResetUptime()
	} else {
		proc.SetUptime()
	}
	proc.SetSysInfo()

	return proc.Status
}

func (proc *Proc) SetStatus(status string) {
	proc.Status.SetStatus(status)
}

func (proc *Proc) SetUptime() {
	proc.Status.SetUptime()
}

func (proc *Proc) ResetUptime() {
	proc.Status.ResetUptime()
}

func (proc *Proc) SetSysInfo() {
	proc.Status.SetSysInfo(proc.process.Pid)
}

func (proc *Proc) Identifier() string {
	return proc.Name
}

func (proc *Proc) ShouldKeepAlive() bool {
	return proc.KeepAlive
}

func (proc *Proc) GetName() string {
	return proc.Name
}