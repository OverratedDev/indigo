package main

import (
  "fmt"
  "os"
)

const version string = "0.1"
const help string = "\nIndigo v" + version + "\n\n\nArgument List:\n\n  new FILENAME          Start a new process\n  end FILENAME          End a process\n  console FILENAME      View a process's input console\n"
const debug bool = true

func debugMsg(msg string) {
  if debug {
    fmt.Println("[DEBUG] " + msg)
  }
}

func main() {
    initFileSystem()

    if len(os.Args) < 2 {
        fmt.Println(help)
        return
    }

    if isWindows() {
      fmt.Println("win");
    }

    arg := os.Args[1]
    fmt.Println(arg)
}
