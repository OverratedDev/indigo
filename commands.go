package main

import (
  "os"
  "strings"
  "fmt"
)

func checkArgs() {
    if len(os.Args) > 1 {
        mainArg := strings.ToLower(os.Args[1])

        if mainArg == "new" {
            if len(os.Args) > 2 {
                if fileExists(os.Args[2]) {
                    fmt.Println("exists")
                } else {
                    fmt.Println("Failed to find file!")
                }
            } else {
              fmt.Println("Invalid arguments! Missing FILENAME")
            }
        } else {
            fmt.Println(help)
            return
        }
    }
}
