package main

import (
	"fmt"
	"time"
	"sync"
)

//Sends and receives are blocking by default
func write(ch chan int) {
	//write to channel
	time.Sleep(1* time.Second)
	for i:=0; i<10; i++ {
		ch <- i
	}
	close(ch) //关闭通道
}

func singleWrite(chn chan<- int) {
	chn <- 122
}

func bufferWrite(bchn chan int) {
	for i := 0; i < 10; i ++ {
		bchn <- i
		fmt.Printf("write done: %d\n", i)
	}
	close(bchn)
}

func process(i int, wg *sync.WaitGroup) {  
    fmt.Println("started Goroutine ", i)
    time.Sleep(2 * time.Second)
    fmt.Printf("Goroutine %d ended\n", i)
    wg.Done()
}

func main() {  
	//declare 双向通道
	ch := make(chan int)
	go write(ch)

	x, ok := <-ch
	if ok == false {
		fmt.Println("Received ", x, ok)
		return
	}

	for data := range ch {
		fmt.Println(data)
	}

	//单向通道
	sendch := make(chan<- int)
	go singleWrite(sendch)
	//ret := <-sendch //invalid operation: <-sendch (receive from send-only type chan<- int)

	/*
		buffered channel，when write the 4th, write is blocked
		no concurrent routine reading from this channel,there will be a deadlock 
		and the program will panic at run time with the following message, 
		fatal error: all goroutines are asleep - deadlock!
	*/
	bchn := make(chan int, 3)
	go bufferWrite(bchn)
	time.Sleep(3 * time.Second)

	for v := range bchn {
		fmt.Println("read value", v,"from ch")
        time.Sleep(2 * time.Second)
	}

	//length is the number of elements in it currently
	//capacity is the buffer size

	//waitGroup is used to wait for all Goroutines finished finish before terminating
	no := 3
    var wg sync.WaitGroup
    for i := 0; i < no; i++ {
		//the WaitGroup's counter is incremented by the value passed to Add
		wg.Add(1)
        go process(i, &wg)
    }
    wg.Wait()
    fmt.Println("All go routines finished executing")

}