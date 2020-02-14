package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {
	//statement
	name := "Hello World"
	fmt.Println(name)
	
	//遍历
	l := len(name)
	for i:= 0; i < l; i++ {
		fmt.Printf("%d, %c\n", name[i], name[i])
	}

	for index, rune := range name {
        fmt.Printf("%c starts at byte %d\n", rune, index)
    }

	//length
	fmt.Printf("length of name: %d", utf8.RuneCountInString(name))

	//unchangable, can't assign to name[0]
	// name[0] = 'a'


}