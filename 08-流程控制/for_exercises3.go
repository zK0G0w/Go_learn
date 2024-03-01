package main

import "fmt"

func main() {
	/*
		输出2-100的素数（只能被1和本身整除）
	*/
	m := 0 //状态码
	for i := 2; i <= 100; i++ {
		for j := 2; j < i; j++ {
			n := i % j
			if n == 0 {
				m = 0
				break
			}
			m = 1
		}
		if m == 1 {
			fmt.Printf("%d为素数\n", i)
		}
	}
}
