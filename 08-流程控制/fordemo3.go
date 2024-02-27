package main

import "fmt"

func main() {
	/*
		循环嵌套：多层循环嵌套在一起
		一：
		*****
		*****
		*****
		*****

		二：
		1x1=1
		2x1=2 2x2=4
		3x1=3 3x2=6 3x3=9
		...
		9x1=9...9x9=81
	*/

	for i := 1; i <= 4; i++ {
		for n := 1; n <= 5; n++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("-----------")

	a := 1
	b := 1
	for a <= 9 {
		for b <= 9 {
			fmt.Printf("%dx%d=%d", a, b, a*b)
			fmt.Println()
			a++
		}
	}
}
