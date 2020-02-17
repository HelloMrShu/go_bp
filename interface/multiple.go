package main

import (
	"fmt"
)

//multiple interfaces
type Run interface {
	run()
}

type Jump interface {
	jump()
}

//type person implement the two interfaces
func (p Person) run()  {
	fmt.Printf("%s could run\n", p.name)
}

func (p Person) jump()  {
	fmt.Printf("%s could jump\n", p.name)
}

type Person struct {
	name string
}

func main() {
	var person = Person {"jackie"}
	person.run()
	person.jump()
}