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
          fmt.Println("Fatal Error: Failed to create " + directory)
          fmt.Println(err)
          return
        }

        debugMsg("Created " + directory)
    }
}
