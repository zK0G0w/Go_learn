package main

import "fmt"

func main() {
	/*
		go语言是静态语言，要求定义，赋值，运算必须类型一致

		类型转换语法格式：Type（Value）

		注意点：兼容类型可以转换

		常数：在有需要的时候，自动转换
		变量：需要手动转换
	*/
	var a int8
	a = 10

	var b int16
	b = int16(a)
	fmt.Println(a, b)

	f1 := 3.54
	var c int
	c = int(f1)
	fmt.Println(f1, c) //3.54, 3

	f1 = float64(a)
	fmt.Println(f1, a)

	//b1 := true
	//a = int8(b1)//cannot convert b1 (variable of type bool) to type int8

	sum := f1 + 100
	fmt.Printf("%T,%f\n", sum, sum)
}
