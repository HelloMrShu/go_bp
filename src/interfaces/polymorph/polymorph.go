package polymorph

import (
	"fmt"
)

type Skills interface {
	Running()
	GetName() string
}

type Student struct {
	Name string
	Age int
}

type Teacher struct {
	Name string
	Salary int
}

//利用method实现多态
func (stu Student) Running() {
	fmt.Printf("student %s is running\n", stu.Name)
}
func (stu Student) GetName() {
	fmt.Printf("student name: %s\n", stu.Name)
}
func (tea Teacher) Running() {
	fmt.Printf("teacher %s is running\n", tea.Name)
}
func (tea Teacher) GetName() {
	fmt.Printf("teacher name: %s\n", tea.Name)
}