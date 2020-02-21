package main

import (  
    "fmt"
	"io/ioutil"
	"os"
    "math/rand"
    "sync"
)

const FILEPATH = "C:/www/github/go_bp/src/files/test.txt"

func read() {
	data, err := ioutil.ReadFile(FILEPATH)
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
	fmt.Println("Contents of file:", string(data))
}

func write() {
	f,err := os.OpenFile(FILEPATH, os.O_WRONLY | os.O_CREATE,0666)
	defer f.Close()
	if err != nil {
        fmt.Println("File reading error", err)
        return
	}
	
	l, err := f.WriteString("Hello World")
    if err != nil {
        fmt.Println(err, l)
        f.Close()
        return
	}
}

func produce(data chan int, wg *sync.WaitGroup) {  
    n := rand.Intn(999)
    data <- n
    wg.Done()
}

func consume(data chan int, done chan bool) {  
    f, err := os.Create("C:/www/github/go_bp/src/files/concurrent")
    if err != nil {
        fmt.Println(err)
        return
    }
    for d := range data {
        _, err = fmt.Fprintln(f, d)
        if err != nil {
            fmt.Println(err)
            f.Close()
            done <- false
            return
        }
    }
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        done <- false
        return
    }
    done <- true
}

func main() {
	read()
	write()
	read()

	data := make(chan int)
    done := make(chan bool)
    wg := sync.WaitGroup{}
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go produce(data, &wg)
    }
    go consume(data, done)
    go func() {
        wg.Wait()
        close(data)
    }()
    d := <-done
    if d == true {
        fmt.Println("File written successfully")
    } else {
        fmt.Println("File writing failed")
    }
}