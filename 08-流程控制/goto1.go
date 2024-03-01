package main

import "fmt"

func main() {
	/*
		goto语句：
		goto label;
		...
		...
		label:statement
	*/

	var a = 10
loop:
	for a < 20 {
		if a == 15 {
			a += 1
			goto loop
		}
		fmt.Printf("a的值为%d\n", a)
		a++
	}
}
