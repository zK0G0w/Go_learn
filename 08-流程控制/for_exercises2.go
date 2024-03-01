package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		水仙花数：三位数：[100，900]
		每个位上的数字的立方和，刚好等于该数字本身，那么就叫做水仙花数

	*/

	a := 0 //个位
	b := 0 //十位
	c := 0 //百位
	for num := 100; num <= 900; num++ {
		a = num / 1 % 10
		b = num / 10 % 10
		c = num / 100 % 10

		if math.Pow(float64(a), 3)+math.Pow(float64(b), 3)+math.Pow(float64(c), 3) == float64(num) {
			fmt.Printf("数字：%d,百位：%d,十位：%d,个位：%d\n", num, c, b, a)
		}

	}
}
