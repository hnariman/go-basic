package main

import "fmt"

func main(){
    defer fmt.Println("one");
    fmt.Println("two");
}
