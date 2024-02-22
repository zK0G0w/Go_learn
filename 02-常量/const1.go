package main

import "fmt"

func main() {
	/*
		常量：
		1.概念：同变量类似，程序执行过程中数值不能改变
		2.语法：
			显式类型定义
			隐式类型定义
	*/
	//1.定义常量
	const PI = 3.14
	const PATH = "www.baidu.com"
	fmt.Println(PI)
	fmt.Println(PATH)

	//2.定义一组常量
	const C1, C2, C3 = 100, 3.14, "haha"
	const (
		MALE    = 0
		FEMALE  = 1
		UNKNOWN = 2
	)
	//3.一组常量中，如果某个常量没有初始值，默认和上一行一致
	const (
		a int = 100
		b
	)
	fmt.Printf("%T,%d\n", a, a)
	fmt.Printf("%T,%d\n", b, b)

	//4.枚举类型：使用常量组作为枚举类型。一组相关数值的数据
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)
}
