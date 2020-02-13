package main

import (
	"fmt"
)

func main()  {
	//声明，大小不可变
	arr := [3]int{1,3} // 等价于var arr = [3]int{1,3}
	fmt.Println(arr)

	//迭代
	l := len(arr) //数组长度
	for i := 0; i < l; i ++ {
		fmt.Printf("index: %d, value: %d\n", i, arr[i])
	}

	for key, val := range arr {
		fmt.Printf("index: %d, value: %d\n", key, val)
	}

	//追加(slice)
	slice := arr[0:]
	slice = append(slice, 100)
	fmt.Println(slice)

	//比较
	arr1 := [3]int{1,2}
	e := arr == arr1
	fmt.Println(e) //false

	//多维数组
	arr2 := [2][2]int{{1,2}, {2,3}}
	fmt.Println(arr2)
}