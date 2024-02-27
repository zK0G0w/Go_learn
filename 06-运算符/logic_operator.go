package main

import "fmt"

func main() {
	/*
			逻辑运算符：
		&&：逻辑与，都为非0为真(一假则假，全真才真)
		||：逻辑或，任何两个为非0为真(一真则真，全假才假)
		！：逻辑非，反转操作数的逻辑状态
	*/

	f1 := true
	f2 := false
	f3 := true
	res1 := f1 && f2
	res2 := f1 || f2
	res3 := f1 && f2 && f3
	res4 := f1 || f2 || f3
	fmt.Printf("res: %t\n", res1) //false
	fmt.Printf("res: %t\n", res2) //true
	fmt.Printf("res: %t\n", res3) //false
	fmt.Printf("res: %t\n", res4) //true

}
