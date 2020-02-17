package main

import (
	"fmt"
)

type Num struct {
	n1 int
	n2 int
}

//接口是包含一组方法定义，表示行为
type Calculate interface {
	add() int
	minute() int
	multiply() int
	divide() int
	mod() int
}

func (num Num) add() int {
	return num.n1 + num.n2
}

func (num Num) minute() int  {
	return num.n1 - num.n2
}

func (num Num) multiply() int  {
	return num.n1 * num.n2
}

func (num Num) divide() int  {
	return num.n1 / num.n2
}

func (num Num) mod() int  {
	return num.n1 % num.n2
}

func main()  {
	num := Num {5,2}
	fmt.Println(num)
	fmt.Printf("add: %d\n", num.add())
	fmt.Printf("minute: %d\n", num.minute())
	fmt.Printf("multiply: %d\n", num.multiply())
	fmt.Printf("divide: %d\n", num.divide())
	fmt.Printf("mod: %d\n", num.mod())
}