package main

import "fmt"

func main() {
	/*
		1.打印58-23
		2.求1-100的和
		3.打印1-100以内，能被3整除但不能被5整除的数字，统计被打印的数字的个数，每行打印5个
	*/
	for a := 58; a >= 23; a-- {
		fmt.Println(a)
	}
	fmt.Println("----------")

	sum := 0
	for b := 1; b <= 100; b++ {
		sum += b
	}
	fmt.Println("1-100的和为：", sum)
	fmt.Println("----------")

	count := 0
	for n := 1; n <= 100; n++ {
		if n%3 == 0 && n%5 != 0 {
			fmt.Print(n, "\t")
			count++
			if count%5 == 0 {
				fmt.Println()
			}
		}
	}
	fmt.Println()
	fmt.Println("满足条件的数有：", count, "个")
}
