package main

import (
    "log"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("./"))
    http.Handle("/", fs)

    log.Println("Listening on http://localhost:8181/index.html")
    err := http.ListenAndServe(":8181", nil)
    if err != nil {
        log.Fatal(err)
    }
}
