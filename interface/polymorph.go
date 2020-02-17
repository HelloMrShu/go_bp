package main

import (
	"fmt"
)

type Skills interface {
	Running()
	GetName() string
}

type Student struct {
	name string
	age int
}

type Teacher struct {
	name string
	salary int
}

//利用method实现多态
func (stu Student) Running() {
	fmt.Printf("student %s is running\n", stu.name)
}
func (stu Student) GetName() {
	fmt.Printf("student name: %s\n", stu.name)
}
func (tea Teacher) Running() {
	fmt.Printf("teacher %s is running\n", tea.name)
}
func (tea Teacher) GetName() {
	fmt.Printf("teacher name: %s\n", tea.name)
}

func main() {
	var stu = Student {"jackie", 25}
	var tea = Teacher {"lee", 5000}

	stu.Running()
	stu.GetName()
	tea.Running()
	tea.GetName()
}