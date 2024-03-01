package main

import "fmt"

func main() {

	arr1 := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	fmt.Println("----------")
	for index, value := range arr1 {
		fmt.Printf("下标是：%d,数值是：%d\n", index, value)
	}

	sum := 0
	for _, v := range arr1 {
		sum += v
	}
	fmt.Println(sum)

	num := 0
	for v := range arr1 {
		num += v
	}
	fmt.Println(num)
}
