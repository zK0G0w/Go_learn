package main

import "fmt"

func main() {
	var num1 int
	num1 = 30
	fmt.Printf("num1的数值是:%d\n", num1)

	var num2 = 20
	fmt.Printf("num2的数值是:%d\n", num2)

	var name = "曾文"
	fmt.Printf("类型是:%T, 数值是:%s\n", name, name)

	sum := 100
	fmt.Println(sum)

	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println(a, b, c)
	//fmt.Println(b)
	//fmt.Println(c)
}
