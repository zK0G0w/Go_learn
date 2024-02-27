package main

import "fmt"

func main() {
	/*
		break和fallthrough语句
		break:可以使用在switch中，也可以用在for循环中
		fallthrough:用于穿透switch分支，被执行的case分支后面紧接的case，无需匹配，执行穿透执行
		fallthrough应该位于case分支最后一行
	*/
	n := 2
	switch n {
	case 1:
		fmt.Println("i m zw")
		fmt.Println("i m zw")
		fmt.Println("i m zw")
	case 2:
		fmt.Println("i m fyc")
		fmt.Println("i m fyc")
		break //强制结束switch
		fmt.Println("i m fyc")
	case 3:
		fmt.Println("i m lhj")
		fmt.Println("i m lhj")
		fmt.Println("i m lhj")
	}

	fmt.Println("-------------")
	m := 2
	switch m {
	case 1:
		fmt.Println("1 season")
	case 2:
		fmt.Println("2 season")
		fallthrough
	case 3:
		fmt.Println("3 season")
	case 4:
		fmt.Println("4 season")
	}
	fmt.Println("main..over...")
}
