package main

import (
    "fmt"
    "os"
    "io/ioutil"
)

type Page struct {
    Header string
    Body string
    Footer string
}

func main() {
    fmt.Println("[Robin] Running... \n")

    if len(os.Args) > 1 {
        file, e := ioutil.ReadFile(os.Args[1])
        if e != nil {
            fmt.Printf("File error: %v\n", e)
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(file))

    } else {
        fmt.Printf("[Robin] Please use a valid path\n")
        os.Exit(1)
    }
}
