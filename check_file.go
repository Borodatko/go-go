package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Enter path: ")
    var filePath string
    fmt.Scanln(&filePath)
    fi, err := os.Stat(filePath)
    exists := !os.IsNotExist(err)
    if exists == true {
        fileSize := fi.Size()
        dateChanges := fi.ModTime()
        dateForma := dateChanges.Format("15:04:05 02-01-2006")
        fmt.Println("File", filePath, "exist.")
        fmt.Println("Size:", fileSize, "b")
        fmt.Println("Modified:", dateForma)
    } else {
        fmt.Println("No such file.")
    }
}
