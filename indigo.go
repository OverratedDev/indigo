package main

import (
  "fmt"
  "os"
)

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

    checkArgs()
}
