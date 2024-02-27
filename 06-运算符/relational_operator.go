package main

import "fmt"

func main() {
	/*
		关系运算符：
		> < == >= <= !=
		结果总是bool类型的:true 和 false
	*/

	a := 3
	b := 5
	//c := 3

	res1 := a > b
	res2 := a < b

	fmt.Printf("%T,%t\n", res1, res1)
	fmt.Printf("%T,%t\n", res2, res2)
}
