package main

import (
	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"fmt"
)	

var (
	app = kingpin.New("indigo", "Process manager designed for application production.")

	start     = app.Command("start", "Create new process.")
	startFile = start.Arg("start file", "file.").Required().String()
	startName = start.Arg("name", "process name").Required().String()

	delete     = app.Command("delete", "Delete a process.")
	deleteName = delete.Arg("name", "process name").Required().String()

	restart     = app.Command("restart", "Restart a process.")
	restartName = restart.Arg("name", "process name").Required().String()

	stop     = app.Command("stop", "Stop a process.")
	stopName = stop.Arg("name", "process name").Required().String()

	list = app.Command("list", "List all processes.")

	status     = app.Command("status", "View a process' status.")
	statusName = status.Arg("name", "process name").Required().String()

	debug = app.Command("debug", "Toggle debug.")
)

func main() {
	checkArgs()
}

func debugMsg(msg string) {
	if indigoDebug {
		fmt.Println("[DEBUG] " + msg)
	}
}
