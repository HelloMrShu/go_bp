package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {  
    time.Sleep(3 * time.Second)
    ch <- "process successful"
}

func server1(ch chan string) {  
    ch <- "from server1"
}

func server2(ch chan string) {  
    ch <- "from server2"

}

func main() {  
    // ch := make(chan string)
    // go process(ch)
    // for {
    //     time.Sleep(1000 * time.Millisecond)
    //     select {
    //     case v := <-ch:
    //         fmt.Println("received value: ", v)
    //         return
    //     default:
    //         fmt.Println("no value received")
    //     }
	// }
	
	output1 := make(chan string)
    output2 := make(chan string)
    go server1(output1)
    go server2(output2)
    time.Sleep(2 * time.Second)
    select {
		case s1 := <-output1:
			fmt.Println(s1)
		case s2 := <-output2:
			fmt.Println(s2)
    }
}