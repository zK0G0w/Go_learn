package main

import "fmt"

func main() {
	/*
		switch 变量名{
		case 数值1：分支1
		case 数值2：分支2
		case 数值3：分支3
		...
		default：
			最后一个分支
		}
	*/

	num := 4
	switch num {
	case 1:
		fmt.Println("一季度")
	case 2:
		fmt.Println("二季度")
	case 3:
		fmt.Println("三季度")
	default:
		fmt.Println("输入错误...")
	}
	fmt.Println("main..over...")
}
