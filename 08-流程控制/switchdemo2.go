package main

import "fmt"

func main() {
	/*
		省略switch后的变量，相当于直接作用在true上
		case后可以同时跟随多个数值
		switch后可以加一条初始化语句，同时变量只能在该switch语句中使用
	*/
	switch {
	case true:
		fmt.Println("true...")
	case false:
		fmt.Println("false...")
	}
}
