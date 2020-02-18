package main

import (
	"fmt"
	"rectangle"
	"circle"
)

func main() {
	length, width, radius := 3,4,5

	r_area := rectangle.Area(length, width)
	fmt.Println("rectangle area is：", r_area)

	c_area := circle.Area(radius)
	fmt.Println("circle area is：", c_area)
}