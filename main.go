package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/pborman/uuid"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)

        log.Print("Hi");

        uuidWithHyphen := uuid.NewRandom()
        fmt.Println(uuidWithHyphen.String())
    })

    http.ListenAndServe(":3000", nil)
}
