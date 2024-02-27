package main

import "fmt"

func main() {
	/*

	 */
	score := 0
	fmt.Println("请输入您的成绩：")
	//fmt.Scanf("%f", &score)
	fmt.Scanln(&score)

	if score >= 60 {
		fmt.Println("您及格了")
	} else {
		fmt.Println("您不及格")
	}

	fmt.Println("main..over...")
}
