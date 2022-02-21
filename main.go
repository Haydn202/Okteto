package main

import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("Starting hello-world server...")
    http.HandleFunc("/", helloServer)
    if err := http.ListenAndServe(":8082", nil); err != nil {
        panic(err)
    }
}

func helloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello okteto2!")
}