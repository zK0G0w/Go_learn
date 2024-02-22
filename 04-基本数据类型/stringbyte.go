package main

import "fmt"

func main() {
	/*
		字符串：
		1.概念：多个byte的集合，理解为一个字符序列
		2.语法：使用双引号 "abc",也可以使用 `A`
		3.编码问题
			计算机本质只识别0和1
			ASCII（美国标准信息交换码）

		4.转义字符：\
			A:有一些字符，有特殊的作用，可以转义为普通的字符，在前面加“\”即可
			B:有一些字符，就是一个普通字符，转义后有特殊作用
				\n换行，\t制表符
	*/

	//1.定义字符串
	var s1 string
	s1 = "曾文"
	fmt.Printf("%T,%s\n", s1, s1)

	s2 := `Hello World`
	fmt.Printf("%T,%s\n", s2, s2)

	//2.""和''区别
	v1 := 'A'
	v2 := "A"
	fmt.Printf("%T\n", v1)        // int32
	fmt.Printf("%T,%s\n", v2, v2) // string

	v3 := '中'
	fmt.Printf("%T,%d,%c,%q\n", v3, v3, v3, v3)

	//3.转义字符
	fmt.Println("\"Hello,world\"")
}
