package main

import (
	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"fmt"
	"os"
)

func checkArgs() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case start.FullCommand():
		startProcess(*startFile, *startName)
	case debug.FullCommand():
		toggleDebug()
	}
}

func toggleDebug() {
	/*if indigoDebug {
			indigoDebug = false
	} else {
			indigoDebug = true
	}*/

	fmt.Println("Debug:", indigoDebug)
}

func startProcess(file string, name string) {

}
