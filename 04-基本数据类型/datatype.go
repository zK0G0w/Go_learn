package main

import "fmt"

func main() {
	/*
		Go语言的数据类型：
		1.基本数据类型：
			布尔类型:bool
				取值：true, false
			数值类型:
				整数: byte = uint8 rune = int32
				浮点数: 小数
					float32, float64
				复数: complex,
			字符串: string
		2.复合数据类型:
			array, slice, map, function, pointer, struct, interface, channel...
	*/
	//1.布尔类型
	var b1 bool
	b1 = true
	fmt.Printf("%T,%t\n", b1, b1)
	b2 := false
	fmt.Printf("%T,%t\n", b2, b2)

	//2.整数
	var i1 int8
	i1 = 100
	fmt.Println(i1)
	//i1 = 200 //报错，int8无法表示200
	var i2 uint8
	i2 = 200
	fmt.Println(i2)

}
