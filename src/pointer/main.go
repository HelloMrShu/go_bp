package main

import (
	"fmt"
)

func main()  {
	//statement, A pointer is a variable which stores the memory address of another variable
	num := 100
	var p *int = &num
	fmt.Printf("type of p: %T\n", p)
	fmt.Printf("address: %v, value: %d\n", p, *p)

	//nil default
	var p1 *int
	fmt.Println(p1) // <nil>

	//new 声明指针
	var p2 = new(int)
	*p2 = 55
	fmt.Printf("address: %v, value: %d\n", p2, *p2)

	//利用指针修改值
	*p2++
	fmt.Printf("address: %v, value: %d\n", p2, *p2)
	
	change(p2)
	fmt.Printf("address: %v, change value: %d\n", p2, *p2)

	addr := getReturn(num)
	fmt.Printf("address: %v, return value: %d\n", addr, *addr)

}

//函数参数-指针
func change(data *int) {
	*data = 111
}

//return pointer
func getReturn(data int) *int {
	data = data + 10
	return &data
}