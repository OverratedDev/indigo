package main

import (
	"os/exec"
	"strings"

	"github.com/vereas/indigo/process"
)

type ProcPreparable interface {
	PrepareBin() ([]byte, error)
	Start() (process.ProcContainer, error)
	getPath() string
	Identifier() string
	getBinPath() string
	getPidPath() string
	getOutPath() string
	getErrPath() string
}

type Preparable struct {
	Name	   string
	SourcePath string
	Cmd        string
	SysFolder  string
	Language   string
	KeepAlive  bool
	Args       []string
}

func (preparable *Preparable) PrepareBin() ([]byte, error) {
	if preparable.SourcePath[len(preparable.SourcePath)-1] == '/' {
		preparable.SourcePath = strings.TrimSuffix(preparable.SourcePath, "/")
	}
	cmd := ""
	cmdArgs := []string{}



	

	preparable.Cmd = preparable.getBinPath()
	return exec.Command(cmd, cmdArgs...).Output()
}

func (preparable *Preparable) Start() (ProcContainer, error) {
	proc := &process.Proc{
		Name:	   preparable.Name,
		Cmd:       preparable.Cmd,
		Args:      preparable.Args,
		Path: 	   preparable.getPath(),
		Pidfile:   preparable.getPidPath(),
		Outfile:   preparable.getOutPath(),
		Errfile:   preparable.getErrPath(),
		KeepAlive: preparable.KeepAlive,
		Status:    &process.ProcStatus{},
	}

	err := proc.Start()
	return proc, err
}

func (preparable *Preparable) Identifier() string {
	return preparable.Name
}

func (preparable *Preparable) getPath() string {
	if preparable.SysFolder[len(preparable.SysFolder)-1] == '/' {
		preparable.SysFolder = strings.TrimSuffix(preparable.SysFolder, "/")
	}
	return preparable.SysFolder + "/" + preparable.Name
}

func (preparable *Preparable) getBinPath() string {
	return preparable.getPath() + "/" + preparable.Name
}

func (preparable *Preparable) getPidPath() string {
	return preparable.getBinPath() + ".pid"
}

func (preparable *Preparable) getOutPath() string {
	return preparable.getBinPath() + ".out"
}

func (preparable *Preparable) getErrPath() string {
	return preparable.getBinPath() + ".err"
}