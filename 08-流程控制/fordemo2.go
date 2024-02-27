package main

import "fmt"

func main() {
	/*
		for循环标准写法中的表达式1和表达式3都可以省略

		如果同时省略三个表达式，则相当于其他语言中的while(true)
	*/
	n := 1
	for n <= 5 {
		fmt.Println(n)
		n++
	}
	fmt.Println("-->", n)
	for {
		fmt.Println("n-->", n)
		n++
	}
}
