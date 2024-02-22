package main

import "fmt"

var a = 1000
var b int = 2000 //定义全局变量
//c := 100 //不能用简短定义的方式来定义全局变量

func main() {
	var num int
	num = 100
	fmt.Printf("num的数值是：%d，地址是：%p\n", num, &num)

	num = 200
	fmt.Printf("num的数值是：%d，地址是：%p\n", num, &num)

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("---------------")
	var m int //整数，默认值是0
	fmt.Println(m)
	var n float64 //0.0 -->0
	fmt.Println(n)
	var s string //""
	fmt.Println(s)
	var s2 []int //切片[]
	fmt.Println(s2)
	fmt.Println(s2 == nil)
}
