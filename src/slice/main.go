package main

import (
	"fmt"
)

func main()  {
	//声明array
	arr := [10]int{1,3,6,9}
	fmt.Println(arr)

	//分片
	slice := arr[3:6]
	
	//类型，array值类型，slice引用类型
	fmt.Printf("\narr type: %T", arr)
	fmt.Printf("\nslice type: %T", slice)

	//长度是具体包含几个元素
	fmt.Printf("\narr type length: %d", len(arr))
	fmt.Printf("\nslice type length: %d", len(slice))

	//slice的容量是从起始地址到底层数据结构的尾部
	fmt.Printf("\narr type capacity: %d", cap(arr))
	fmt.Printf("\nslice type capacity: %d\n\n", cap(slice))

	slice = append(slice, 55)

	fmt.Println(slice)

}