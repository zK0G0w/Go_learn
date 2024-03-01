package main

import "fmt"

func main() {
	/*
		数据类型：
			基本类型：int,float,string,bool...
			复合类型：array,slice,map,function,pointer,channel...

			数组的数据类型：[size]type
	*/

	num := 10
	fmt.Printf("%T\n", num)

	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [3]float64{2.15, 3.18, 6.19}
	arr3 := [4]int{5, 6, 7, 8}
	arr4 := [2]string{"hello", "world"}

	fmt.Printf("%T\n", arr1) //[4]int
	fmt.Printf("%T\n", arr2) //[3]float64
	fmt.Printf("%T\n", arr3) //[4]int
	fmt.Printf("%T\n", arr4) //[2]string

}
