package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		生成随机数：
		伪随机数，根据一定的算法公式算出来的
	*/

	num1 := rand.Int()
	fmt.Println(num1)

	for i := 0; i < 10; i++ {
		num := rand.Intn(10)
		fmt.Println(num)
	}
	for i := 0; i < 100; i++ {
		t1 := time.Now()
		timeStamp1 := t1.Unix()
		fmt.Println(timeStamp1)
		time.Sleep(1 * time.Second)
	}
}
