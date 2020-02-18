package main

import (
	"fmt"
	"log"
)

func main()  {
	//statement
	user := make(map[string]int)
	scores := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }

	fmt.Printf("Type: %T\n", user)
	fmt.Println(user)
	fmt.Println(scores)
	
	//add
	user["jim"] = 100
	
	//get
	key := "tom"
	val, ok := user[key]
	if ok == true {
		fmt.Println("get value: ", val)
	} else {
		log.Printf("%s not found", key)
	}

	//iterate
	for key, val := range user {
		fmt.Printf("key is %s, value is %d", key, val)
	}

	//delete
	delete(user, "steve")

	//length
	fmt.Printf("\nuser map length: %d", len(user))

	//compare
	if user == scores {
		fmt.Println(true)
	}

	//map can only be compared to nil, you should check each element in both maps
	//flag := user == scores //error

}