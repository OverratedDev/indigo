package main

import (
	"os"
	"syscall"

	"github.com/sevlyar/go-daemon"
)

func hasDaemonStarted(ctx *daemon.Context) (bool, *os.Process, error) {
	d, err := ctx.Search()

	if err != nil {
		return false, d, err
	}

	if err := d.Signal(syscall.Signal(0)); err != nil {
		return false, d, err
	}

	return true, d, nil
}
