package multiple

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
func (p Person) Run()  {
	fmt.Printf("%s could run\n", p.Name)
}

func (p Person) Jump()  {
	fmt.Printf("%s could jump\n", p.Name)
}

type Person struct {
	Name string
}