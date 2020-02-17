package main

/*
	A method is just a function with a special receiver type between the func keyword and the method name. 
	The receiver can either be a struct type or non-struct type.
	two method have same name with different type is ok, but not for func.
	there is no class in go, but methods allow a logical grouping of behavior 
	related to a type similar to classes
*/

import (
	"fmt"
)

type User struct {  
    name  string
    age   int
}

//function
func displayFunc(u User) {  
    fmt.Printf("function: age of %s is %d\n", u.name, u.age)
}

//method
func (u User) displayMethod() {  
    fmt.Printf("method: age of %s is %d\n", u.name, u.age)
}

func main()  {
	user := User{"jackie", 25}
	displayFunc(user)
	user.displayMethod()
}