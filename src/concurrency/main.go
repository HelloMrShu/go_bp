package main

import (  
	"fmt"
	"time"
)

//goroutine, concurrency
//just invode the function without waiting for return
func hello() {
	fmt.Println("\nHello world goroutine")
	
}

func display() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}
}
func main() {  
	go hello()
	go display()
	time.Sleep(1 * time.Second)
    fmt.Println("\nmain function")
}