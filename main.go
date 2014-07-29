package main

import (
	"fmt"
	"net/http"
)

func main() {
    fmt.Println("[Robin] building server...")

    http.Handle("/", http.FileServer(http.Dir("./")))
    http.ListenAndServe(":8123", nil)
}

