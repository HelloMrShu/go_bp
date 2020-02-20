package main

import (
    "fmt"
    "sync"
)

var x  = 0

// use mutex to solve race condition problems
func increment(wg *sync.WaitGroup, mu *sync.Mutex) {
    mu.Lock()
    x = x + 1
    mu.Unlock()
    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    var mu sync.Mutex

    for i := 0; i < 1000; i++ {
        wg.Add(1)        
        go increment(&wg, &mu)
    }
    wg.Wait()
    fmt.Println("final value of x", x)
}