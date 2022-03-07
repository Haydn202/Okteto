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

var visNumber int8 = 0

func helloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hit endpoint")
    visNumber++
    visString := fmt.Sprintf("%d", visNumber)
    fmt.Fprint(w, "Hello Okteto"+ visString +"!")
}