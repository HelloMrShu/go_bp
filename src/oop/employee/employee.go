package employee

import (
	"fmt"
)

//大写：访问权限public, 小写包内访问
type employee struct {
	name   string
	age int
}

func New(name string, age int) employee {
	e := employee{name, age}
	return e
}

func (e employee) Display() {
    fmt.Printf("%s is %d years old!", e.name, e.age)
}