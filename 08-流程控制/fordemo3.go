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

	/*a := 1   //前面的数
	b := 1   //后面的数
	col := 1 //列数
	row := 1 //行数
	for row <= 9 {
		for col <= 9 {
			for m := 1; m <= row; m++ {
				fmt.Printf("%dx%d=%d ", a, b, a*b)
				b++
			}
			b = 1
			fmt.Println()
			row++
			col++
			a++
		}
	}
	我的方法太笨了我靠
	*/

	for m := 1; m < 10; m++ {
		for j := 1; j <= m; j++ {
			fmt.Printf("%dx%d=%d\t", j, m, m*j)
		}
		fmt.Println()
	}
}
