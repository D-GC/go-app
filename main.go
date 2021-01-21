package main

import (
    "fmt"
    "net/http"
    "log"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
        log.fmt("Hello");
        mgo.Hello;
    })

    http.ListenAndServe(":3000", nil)
}
