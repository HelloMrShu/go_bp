package main

import (
    "fmt"
    "time"
    // "runtime/debug"
)

//defer panic recover 处理异常
func recovery() {  
    if r := recover(); r != nil {
        fmt.Println("recovered:", r)
        // debug.PrintStack()
    }
}

func a() {  
    defer recovery()
    fmt.Println("Inside A")
    b()
    time.Sleep(1 * time.Second)
}

func b() {  
    fmt.Println("Inside B")
    panic("oh! B panicked")
}

func c() {
    defer recovery()
    fmt.Println("Inside C")
    panic("C is panicking")
}

func main() {  
    a()
    fmt.Println("A normally returned from main")

    fmt.Printf("\n")

    c()
    fmt.Println("C normally returned from main")
}