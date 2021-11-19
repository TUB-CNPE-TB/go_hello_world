package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    var jsonData = []byte(`{
        "message": "Hello World!"
    }`)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonData)
    })

    fmt.Printf("Starting server at port 9090\n")
    if err := http.ListenAndServe(":9090", nil); err != nil {
        log.Fatal(err)
    }
}
