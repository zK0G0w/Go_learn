package main

import "fmt"

func main() {
	var num1 int
	num1 = 100
	num1 = 200
	fmt.Println(num1)

	//step1 创建数组
	var arr1 [4]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4

	for i := 0; i <= 3; i++ {
		fmt.Println(arr1[i])
	}
	fmt.Println("数组的长度：", len(arr1)) //实际存储的数量
	fmt.Println("数组的容量：", cap(arr1)) //能够存储的数量

	var a [4]int
	fmt.Println(a)

	var b = [4]int{1, 2, 3, 4}
	fmt.Println(b)

	var c = [5]int{1, 2, 3}
	fmt.Println(c)

	var d = [5]int{1: 1, 3: 4}
	fmt.Println(d)

	e := [...]int{1: 2, 4: 4}
	fmt.Println(e)

}
