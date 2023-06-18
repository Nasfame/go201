package main

import (  
    "fmt"
    "net/http"
)

func main() {  
    fmt.Println("starting server")
    err := http.ListenAndServe(":9090", http.FileServer(http.Dir("../")))
    if err != nil {
        fmt.Println("Failed to start server", err)
        return
    }
}