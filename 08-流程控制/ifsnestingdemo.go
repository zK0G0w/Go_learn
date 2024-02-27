package main

import "fmt"

func main() {
	/*
	 */

	sex := "泰国"

	if sex == "男" {
		fmt.Println("请去男厕所")
	} else if sex == "女" {
		fmt.Println("请去女厕所")
	} else {
		fmt.Println("请使用第三性别厕所")
	}

	fmt.Println("main..over...")
}
