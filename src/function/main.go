package main

import (
	"fmt"
)

func findv1(num int, nums []int) {  
    fmt.Printf("type of v1 nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "founded at index", i)
			found = true
			break
        }
    }
    if !found {
        fmt.Println(num, "not founded")
    }
    fmt.Printf("\n")
}

//variadic function
func findv2(num int, nums ...int) {  
    fmt.Printf("type of v2 nums is %T\n", nums)
    found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "founded at index", i)
			found = true
			break
        }
    }
    if !found {
        fmt.Println(num, "not founded")
    }
	fmt.Printf("\n")
}

func main() {
	nums := []int{89, 90, 95,45}

	findv1(95, []int{89, 90, 95, 45})
    findv2(45, nums...)
}