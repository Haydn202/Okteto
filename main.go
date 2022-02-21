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

var x int8 = 0

func helloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hit endpoint")
    x++
    s := fmt.Sprintf("%d", x)
    fmt.Fprint(w, "Hello visitor"+ s +"!")
}