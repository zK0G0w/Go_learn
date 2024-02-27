package main

import "fmt"

func main() {
	/*
			位运算符：将数值转为二进制后，按位操作
		按位&:对应位的值如果都为1才为1，有一个位0就为0
		按位|:对对应位的值如果都是0才为0，有一个为1就为1
		异或^:
			二元:a^b
				对应位的值不同置为1，相同为0
			一元：^a
				按位取反
		位清空&^:
			对于 a &^ b
				对于b上的每个数值
				如果为0，则取a对应位上的数值
				如果为1，则结果位取0

		位移运算符：
		<<：按位左移，a << b ,把a转换为二进制，左移b位，即放大2的b次方倍
		>>：按位右移，a >> b ,把a转换为二进制，右移b位，即缩小2的b次方倍
	*/

	a := 60
	b := 13
	/*
		a:60, 0011 1100
		b:13, 0000 1101
		&:12, 0000 1100
		|:61, 0011 1101
		^:49, 0011 0001
		&^48, 0011 0000
	*/
	fmt.Printf("a:%d, %b\n", a, a)
	fmt.Printf("b:%d, %b\n", b, b)
	fmt.Printf("a&b:%d, %b\n", a&b, a&b)
	fmt.Printf("a|b:%d, %b\n", a|b, a|b)
	fmt.Printf("a^b:%d, %b\n", a^b, a^b)
	fmt.Printf("a&^b:%d, %b\n", a&^b, a&^b)
	fmt.Printf("^a:%d, %b\n", ^a, ^a)

	a1 := a >> 2
	a2 := a << 2
	fmt.Printf("a >> 2:%d,%b\n", a1, a1)
	fmt.Printf("a << 2:%d,%b\n", a2, a2)

}
