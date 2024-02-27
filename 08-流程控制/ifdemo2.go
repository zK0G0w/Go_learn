package main

import "fmt"

func main() {
	/*
		if语句的其他写法，可以在if语句中定义变量，作用域为该if语句
	*/

	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}
