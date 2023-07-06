//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 数组的核心就是数组的下标
	 */
	arr := [5]int{1, 3, 7, 21, 42}
	fmt.Println(arr)

	for index, value := range arr {
		fmt.Printf("%d -> %d\n", index, value)
	}

	//只要数值，不要下标
	for _, value := range arr {
		fmt.Printf("[%d]\n", value)
	}
}
