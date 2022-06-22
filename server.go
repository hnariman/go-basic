package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request){
    fmt.Fprint(w, "hello\n")
}

func main(){
    http.HandleFunc("/", hello)

    http.ListenAndServe(":8090", nil)
}
