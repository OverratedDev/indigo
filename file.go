package main

import (
  "os"
  "fmt"
)

const directory string = "/etc/.indigo"

func initFileSystem() {
    if _, err := os.Stat(directory); os.IsNotExist(err) {
        err := os.MkdirAll(directory, 0711)

        if err != nil {
            fmt.Println(err)
            return
        }

        debugMsg("Created " + directory)

        createSessionsFile()
    }
}

func fileExists(file string) bool {
    if _, err := os.Stat(file); os.IsNotExist(err) {
        if err != nil {
            fmt.Println(err)
            return false
        }

        return false
    } else {
        return true
    }
}

func createSessionsFile() {
    file, err := os.Create(directory + "/.sessions")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    debugMsg("Created " + directory + "/.sessions")

    fmt.Fprintf(file, "# SESSION-NAME -- SCRIPT -- START-DATE")
}
