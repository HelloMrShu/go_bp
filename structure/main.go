package main

import (
	"fmt"
)

func main()  {
	//declare
	type Address struct {  
		city string
	}

	type User struct {
		name string
		age int
		addr Address //nested struct
	}

	u1 := User {
		name: "jim",
		age: 16,
		addr: Address {
			city: "beijing",
		},
	}

	u2 := User{"tom", 20, Address{"shanghai"}}

	fmt.Println(u1)
	fmt.Println(u2)

	//default value without initialized
	var u3 User
	fmt.Println(u3)

	//specify values for some fields
	var u4 User
	u4.name = "tinker"
	fmt.Println(u4)

	//pointer to struct
	var p = &User{"lily", 25, Address{"guangzhou"}}
	p.age = 33
	fmt.Println(*p)
}