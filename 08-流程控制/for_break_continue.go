package main

import "fmt"

func main() {
	/*
			循环控制语句
			break：彻底结束循环
			continue：结束了某一次循环，下次循环继续
		可以通过给循环名标签名，来要求break和continue针对的循环体，否则就是结束最近的里层循环
	*/

	for i := 1; i <= 10; i++ {
		fmt.Print(i)
		if i > 5 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("-------------")
in:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if j == 2 {
				break in
			}
			fmt.Printf("i:%d,j:%d\n", i, j)
		}
	}
}
