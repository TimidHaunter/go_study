//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	/*
	 * 数组排序，让数组元素有一定的顺序
	 * 冒泡，go没有自己的排序函数吗？
	 * 升序 [2, 4, 5, 6]
	 * 降序 [6, 5, 4, 2]
	 */
	var arr = [4]int{6, 4, 5, 2}
	fmt.Println(arr)
	for j := 0; j < (len(arr) - 1); j++ {
		if arr[j] > arr[j+1] {
			//交换两个元素
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
	fmt.Println(arr)

	for i := 1; i < len(arr); i++ {
		for j := 0; j < (len(arr) - i); j++ {
			if arr[j] > arr[j+1] {
				//交换两个元素
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
