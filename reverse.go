package main

import (
    "fmt"
    s "strings"
)

func main(){
    text := "this is string"
    fmt.Println(text)
    words := s.Split("esrever era ylno sdrow", " ")


    slice:= make([]string, 0)

    for i:=len(text)-1 ; i >= 0; {
        fmt.Print(string(text[i]))
        slice = append(slice, string(text[i]))
        i--
    }

    fmt.Println("")
    fmt.Println(slice)

    sl := ""
    for _, word := range words {
        sl = sl + " "
        for i := range word {
           sl  = sl + string(word[len(word)-1-i])
        }
    }
    fmt.Println(sl)

}

