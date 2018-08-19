package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        hostname, err := os.Hostname()
        if err != nil {
            fmt.Fprintf(w, "Could not get hostname")
            return
        }

        fmt.Fprintf(w, hostname)
    })


    log.Fatal(http.ListenAndServe(":8080", nil))
}
